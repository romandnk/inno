package httphandler

import (
	_ "authservice/docs"
	"authservice/internal/domain"
	"authservice/internal/service"
	"errors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

//	@title			Auth service API
//	@version		1.0
//	@description	This is an auth service.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/swagger/*", httpSwagger.Handler())

	router.Handle("/sign_up", CORS(LogUser(http.HandlerFunc(SignUp))))
	router.Handle("/sign_in", CORS(LogUser(http.HandlerFunc(SignIn))))

	router.Handle("/get_user_info", CORS(Auth(LogUser(http.HandlerFunc(GetUserInfo)))))
	router.Handle("/set_user_info", CORS(Auth(LogUser(http.HandlerFunc(SetUserInfo)))))
	router.Handle("/change_psw", CORS(Auth(LogUser(http.HandlerFunc(ChangePsw)))))
	// admin handlers
	router.Handle("/admin/get_user_info", CORS(Auth(isAdmin(LogUser(http.HandlerFunc(AdminGetUserInfo))))))
	router.Handle("/admin/set_user_block_status", CORS(Auth(isAdmin(LogUser(http.HandlerFunc(AdminSetUserBlockStatus))))))
	router.Handle("/admin/set_user_role", CORS(Auth(isAdmin(LogUser(http.HandlerFunc(AdminSetUserRole))))))
	router.Handle("/admin/reset_user_password", CORS(Auth(isAdmin(LogUser(http.HandlerFunc(AdminResetUserPassword))))))

	router.Handle("/v2/get_user_info", CORS(Auth(LogUser(http.HandlerFunc(GetUserInfoV2)))))

	return router
}

// GetUserInfoV2 godoc
//	@Summary		Get User Information (Version 2)
//	@Description	Retrieves user information based on the role of the requesting user. Admins can get full info for any user, while regular users can only access their own short info.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string								true	"User ID to fetch information for"
//	@Param			User-ID	header		string								true	"Author User ID (the requester)"
//	@Success		200		{object}	HTTPResponse{data=domain.User}		"Full information for admins"
//	@Success		200		{object}	HTTPResponse{data=domain.UserInfo}	"Short information for regular users"
//	@Failure		400		{object}	HTTPResponse						"Invalid user ID format"
//	@Failure		403		{object}	HTTPResponse						"Forbidden: Regular user cannot access other user's info"
//	@Failure		404		{object}	HTTPResponse						"User not found"
//	@Failure		500		{object}	HTTPResponse						"Internal server error"
//	@Router			/v2/get_user_info [get]
func GetUserInfoV2(resp http.ResponseWriter, req *http.Request) {

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

	authorUserID, _ := primitive.ObjectIDFromHex(req.Header.Get(HeaderUserID))
	authorFullInfo, err := service.GetUserFullInfo(authorUserID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		respBody.SetError(err)
		return
	}

	switch authorFullInfo.Role {
	case domain.UserRoleAdmin:
		info, err := service.GetUserFullInfo(userID)
		if err != nil {
			resp.WriteHeader(http.StatusNotFound)
			respBody.SetError(err)
			return
		}
		respBody.SetData(info)
		return
	case domain.UserRoleDefault:
		if authorUserID != userID {
			resp.WriteHeader(http.StatusForbidden)
			respBody.SetError(errors.New("user is not admin to get info about someone else"))
			return
		}
		info, err := service.GetUserShortInfo(userID)
		if err != nil {
			resp.WriteHeader(http.StatusNotFound)
			respBody.SetError(err)
		}

		respBody.SetData(info)
	}
}
