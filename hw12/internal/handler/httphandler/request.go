package httphandler

import (
	"authservice/internal/domain"
)

type SetUserInfoReq struct {
	Name string `json:"name"`
}

type ChangePswReq struct {
	Password string `json:"password"`
}

func (r SetUserInfoReq) IsValid() bool {
	return r.Name != ""
}

func (r ChangePswReq) IsValid() bool {
	return r.Password != ""
}

type SetUserBlockStatusReq struct {
	UserID         string `json:"user_id"`
	SetBlockStatus bool   `json:"set_block_status"`
}

type SetUserRoleReq struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func (r SetUserRoleReq) IsValid() bool {
	if r.Role != domain.UserRoleDefault && r.Role != domain.UserRoleAdmin {
		return false
	}
	return true
}

type ResetUserPasswordReq struct {
	UserID string `json:"user_id"`
}
