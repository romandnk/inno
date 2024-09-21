package httphandler

import (
	"authservice/internal/service"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AdminGetUserInfo(resp http.ResponseWriter, req *http.Request) {

	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	id := req.URL.Query().Get("user_id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		respBody.SetError(errors.New("invalid input"))
		return
	}

	info, err := service.GetUserFullInfo(userID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
	}

	respBody.SetData(info)
}

func AdminSetUserBlockStatus(resp http.ResponseWriter, req *http.Request) {
	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		respBody.SetError(fmt.Errorf("only method %s is allowed", http.MethodPost))
		return
	}

	var input SetUserBlockStatusReq

	if err := readBody(req, &input); err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		respBody.SetError(err)
		return
	}

	userID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		respBody.SetError(errors.New("invalid input"))
		return
	}

	err = service.SetUserBlockStatus(userID, input.SetBlockStatus)
	switch {
	case service.IsErrUserNotFound(err):
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
		return
	case service.IsErrUserAlreadyBlocked(err):
		resp.WriteHeader(http.StatusConflict)
		respBody.SetError(err)
		return
	case service.IsErrUserAlreadyUnblocked(err):
		resp.WriteHeader(http.StatusConflict)
		respBody.SetError(err)
		return
	}
}

func AdminSetUserRole(resp http.ResponseWriter, req *http.Request) {
	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		respBody.SetError(fmt.Errorf("only method %s is allowed", http.MethodPost))
		return
	}

	var input SetUserRoleReq

	if err := readBody(req, &input); err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		respBody.SetError(err)
		return
	}

	if !input.IsValid() {
		resp.WriteHeader(http.StatusBadRequest)
		respBody.SetError(errors.New("invalid input"))
		return
	}

	userID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		respBody.SetError(errors.New("invalid input"))
		return
	}

	err = service.SetUserRole(userID, input.Role)
	switch {
	case service.IsErrUserNotFound(err):
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
		return
	case service.IsErrRoleAlreadySet(err):
		resp.WriteHeader(http.StatusConflict)
		respBody.SetError(err)
		return
	}
}

func AdminResetUserPassword(resp http.ResponseWriter, req *http.Request) {
	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	if req.Method != http.MethodPost {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		respBody.SetError(fmt.Errorf("only method %s is allowed", http.MethodPost))
		return
	}

	var input ResetUserPasswordReq

	if err := readBody(req, &input); err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		respBody.SetError(err)
		return
	}

	userID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		respBody.SetError(errors.New("invalid input"))
		return
	}

	newPassword, err := service.ResetUserPassword(userID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
		return
	}

	respBody.SetData(map[string]string{"new_password": newPassword})
}
