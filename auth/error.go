package auth

import "github.com/pkg/errors"

var (
	ErrUserNotFound       = errors.New("user not founds")
	ErrInvalidAccessToken = errors.New("invalid access token")
)
