package service

import "errors"

var (
	ErrUserAlreadyBlocked   = errors.New("user already blocked")
	ErrUserAlreadyUnblocked = errors.New("user already unblocked")
)

func IsErrUserAlreadyBlocked(err error) bool {
	return errors.Is(err, ErrUserAlreadyBlocked)
}

func IsErrUserAlreadyUnblocked(err error) bool {
	return errors.Is(err, ErrUserAlreadyUnblocked)
}

var ErrUserNotFound = errors.New("user not found")

func IsErrUserNotFound(err error) bool {
	return errors.Is(err, ErrUserNotFound)
}

var ErrRoleAlreadySet = errors.New("role already set")

func IsErrRoleAlreadySet(err error) bool {
	return errors.Is(err, ErrRoleAlreadySet)
}
