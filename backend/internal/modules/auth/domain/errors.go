package domain

import "errors"

var (
	ErrInvalidPayload      = errors.New("invalid payload")
	ErrUsernameTaken       = errors.New("username taken")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrUserNotFound        = errors.New("user not found")
	ErrHashingFailed       = errors.New("failed to hash password")
	ErrTokenIssueFailed    = errors.New("failed to issue token")
)
