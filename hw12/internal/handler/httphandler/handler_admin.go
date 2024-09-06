package httphandler

import (
	"authservice/internal/service"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// AdminGetUserInfo godoc
//	@Summary		Get Full User Information
//	@Description	Allows an admin to retrieve the full information of any user by user ID.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string							true	"User ID to fetch full information for"
//	@Success		200		{object}	HTTPResponse{data=domain.User}	"Successfully retrieved full user information"
//	@Failure		400		{object}	HTTPResponse					"Invalid user ID format"
//	@Failure		404		{object}	HTTPResponse					"User not found"
//	@Router			/admin/get_user_info [get]
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

// AdminSetUserBlockStatus godoc
//	@Summary		Set User Block/Unblock Status
//	@Description	Allows an admin to block or unblock a user by setting the user's block status.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			input	body		SetUserBlockStatusReq	true	"User block/unblock status request"
//	@Success		200		{object}	HTTPResponse			"Successfully changed the user's block status"
//	@Failure		400		{object}	HTTPResponse			"Invalid input data"
//	@Failure		404		{object}	HTTPResponse			"User not found"
//	@Failure		409		{object}	HTTPResponse			"User already in desired block/unblock status"
//	@Failure		422		{object}	HTTPResponse			"Request body validation error"
//	@Failure		405		{object}	HTTPResponse			"Method not allowed"
//	@Router			/admin/set_user_block_status [post]
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

// AdminSetUserRole godoc
//	@Summary		Set User Role
//	@Description	Allows an admin to set the role of a user.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			input	body		SetUserRoleReq	true	"User role change request"
//	@Success		200		{object}	HTTPResponse	"Successfully changed the user's role"
//	@Failure		400		{object}	HTTPResponse	"Invalid input data"
//	@Failure		404		{object}	HTTPResponse	"User not found"
//	@Failure		409		{object}	HTTPResponse	"User already has the specified role"
//	@Failure		422		{object}	HTTPResponse	"Request body validation error"
//	@Failure		405		{object}	HTTPResponse	"Method not allowed"
//	@Router			/admin/set_user_role [post]
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

// AdminResetUserPassword godoc
//	@Summary		Reset User Password
//	@Description	Allows an admin to reset a user's password and retrieve the new password.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			input	body		ResetUserPasswordReq					true	"User password reset request"
//	@Success		200		{object}	HTTPResponse{data=map[string]string}	"Successfully reset the user's password"
//	@Failure		400		{object}	HTTPResponse							"Invalid input data"
//	@Failure		404		{object}	HTTPResponse							"User not found"
//	@Failure		405		{object}	HTTPResponse							"Method not allowed"
//	@Failure		422		{object}	HTTPResponse							"Request body validation error"
//	@Router			/admin/reset_user_password [post]
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
