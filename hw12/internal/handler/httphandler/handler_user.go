package httphandler

import (
	"authservice/internal/domain"
	"authservice/internal/service"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

// SignUp godoc
//
//	@Summary		User Registration
//	@Description	Creates a new user account with a login and password.
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.LoginPassword				true	"User login and password"
//	@Success		200		{object}	HTTPResponse{data=domain.UserToken}	"Successful registration, JWT token"
//	@Failure		400		{object}	HTTPResponse						"Invalid registration data"
//	@Failure		409		{object}	HTTPResponse						"User input data error"
//	@Failure		422		{object}	HTTPResponse						"Request body validation error"
//	@Router			/sign_up [post]
func SignUp(resp http.ResponseWriter, req *http.Request) {

	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	var input domain.LoginPassword
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

	userToken, err := service.SignUp(&input)
	if err != nil {
		resp.WriteHeader(http.StatusConflict)
		respBody.SetError(err)
		return
	}

	respBody.SetData(userToken)
}

// SignIn godoc
//
//	@Summary		User Login
//	@Description	Authenticates a user using login and password, returning a JWT token upon success.
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.LoginPassword				true	"User login and password"
//	@Success		200		{object}	HTTPResponse{data=domain.UserToken}	"Successful login, JWT token"
//	@Failure		400		{object}	HTTPResponse						"Invalid login or password"
//	@Failure		404		{object}	HTTPResponse						"User not found"
//	@Failure		422		{object}	HTTPResponse						"Request body validation error"
//	@Router			/sign_in [post]
func SignIn(resp http.ResponseWriter, req *http.Request) {

	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	var input domain.LoginPassword
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

	userToken, err := service.SignIn(&input)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
		return
	}

	respBody.SetData(userToken)
}

// GetUserInfo godoc
//
//	@Summary		Get User Information
//	@Description	Retrieves short information about a user based on the provided user ID.
//	@Tags			user
//	@Produce		json
//	@Param			User-ID	header		string								true	"User ID"
//	@Success		200		{object}	HTTPResponse{data=domain.UserInfo}	"Successfully retrieved user information"
//	@Failure		404		{object}	HTTPResponse						"User not found"
//	@Router			/get_user_info [get]
func GetUserInfo(resp http.ResponseWriter, req *http.Request) {

	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	userID, _ := primitive.ObjectIDFromHex(req.Header.Get(HeaderUserID))

	info, err := service.GetUserShortInfo(userID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
	}

	respBody.SetData(info)
}

// SetUserInfo godoc
//
//	@Summary		Update User Information
//	@Description	Updates user information such as name based on the provided user ID.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			User-ID	header		string			true	"User ID"
//	@Param			input	body		SetUserInfoReq	true	"User information to update"
//	@Success		200		{object}	HTTPResponse	"Successfully updated user information"
//	@Failure		400		{object}	HTTPResponse	"Invalid input data"
//	@Failure		404		{object}	HTTPResponse	"User not found"
//	@Failure		422		{object}	HTTPResponse	"Request body validation error"
//	@Router			/set_user_info [put]
func SetUserInfo(resp http.ResponseWriter, req *http.Request) {

	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	var input SetUserInfoReq

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

	userID, _ := primitive.ObjectIDFromHex(req.Header.Get(HeaderUserID))

	if err := service.SetUserInfo(&domain.UserInfo{
		ID:   userID,
		Name: input.Name,
	}); err != nil {
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
		return
	}
}

// ChangePsw godoc
//	@Summary		Change User Password
//	@Description	Changes the user's password based on the provided user ID and new password.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			User-ID	header		string			true	"User ID"
//	@Param			input	body		ChangePswReq	true	"New password information"
//	@Success		200		{object}	HTTPResponse	"Password successfully changed"
//	@Failure		400		{object}	HTTPResponse	"Invalid input data"
//	@Failure		404		{object}	HTTPResponse	"User not found"
//	@Failure		422		{object}	HTTPResponse	"Request body validation error"
//	@Router			/change_psw [put]
func ChangePsw(resp http.ResponseWriter, req *http.Request) {

	respBody := &HTTPResponse{}
	defer func() {
		resp.Write(respBody.Marshall())
	}()

	var input ChangePswReq

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

	userID, _ := primitive.ObjectIDFromHex(req.Header.Get(HeaderUserID))
	err := service.ChangePsw(&domain.UserPassword{
		ID:       userID,
		Password: input.Password,
	})
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		respBody.SetError(err)
		return
	}
}

func readBody(req *http.Request, s any) error {

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, s)
}
