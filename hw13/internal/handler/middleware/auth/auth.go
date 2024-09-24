package auth

import (
	"crud/internal/pkg/authclient"
	"github.com/valyala/fasthttp"
	"log"
)

var (
	UserID struct{}
	Role   struct{}
)

func Auth(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := ctx.Request.Header.Peek(fasthttp.HeaderAuthorization)
		if string(token) == "" {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			log.Println("Get request", string(ctx.Method()), string(token), "error", fasthttp.StatusUnauthorized)
			return
		}

		response, isValid := authclient.ValidateToken(string(token))
		if !isValid {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			log.Println("Get request", string(ctx.Method()), string(token), "error", fasthttp.StatusUnauthorized)
			return
		}

		ctx.SetUserValue(UserID, response.UserId)
		ctx.SetUserValue(Role, response.Role)

		requestHandler(ctx)
	}
}
