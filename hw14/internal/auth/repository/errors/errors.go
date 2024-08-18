package storageerrors

import "errors"

var ErrNicknameAlreadyExists = errors.New("nickname already exists")

var ErrTokenNotFound = errors.New("token not found")
