package handler

import (
	"crud/internal/domain"
	"crud/internal/handler/middleware/auth"
	"crud/internal/service"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)

func ServerHandler(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowOrigin, "*")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodPost)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodGet)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, fasthttp.MethodDelete)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, fasthttp.HeaderContentType)
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, fasthttp.HeaderAuthorization)

	if ctx.IsOptions() {
		return
	}

	switch {
	case ctx.IsGet():
		GetHandler(ctx)
	case ctx.IsDelete():
		auth.Auth(DeleteHandler)
	case ctx.IsPost():
		auth.Auth(PostHandler)
	}
}

func GetHandler(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")
	if len(id) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	rec, err := service.Get(string(id))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	marshal, err := json.Marshal(rec)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if _, err = ctx.Write(marshal); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func DeleteHandler(ctx *fasthttp.RequestCtx) {
	userID, ok := getUserIDFromCtx(ctx)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	userRole, ok := getRoleFromCtx(ctx)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	id := ctx.QueryArgs().Peek("id")
	if len(id) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	recipe, err := service.Get(string(id))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	if userRole != auth.AdminRole && recipe.UserID != userID {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}

	if err := service.Delete(string(id)); err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func PostHandler(ctx *fasthttp.RequestCtx) {
	userID, ok := getUserIDFromCtx(ctx)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	var rec domain.Recipe
	log.Println(string(ctx.PostBody()))
	if err := json.Unmarshal(ctx.PostBody(), &rec); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	rec.UserID = userID
	if err := service.AddOrUpd(&rec); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	resp := IdResponse{ID: rec.ID}

	marshal, err := json.Marshal(resp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if _, err = ctx.Write(marshal); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func getUserIDFromCtx(ctx *fasthttp.RequestCtx) (string, bool) {
	userID, ok := ctx.UserValue(auth.UserID).(string)
	return userID, ok
}

func getRoleFromCtx(ctx *fasthttp.RequestCtx) (string, bool) {
	role, ok := ctx.UserValue(auth.Role).(string)
	return role, ok
}
