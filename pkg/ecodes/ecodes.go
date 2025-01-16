package ecodes

import "errors"

var (
	ErrUnauthorized   = errors.New("unauthorized request")
	ErrEmailExist     = errors.New("email already exist")
	ErrInvalidReq     = errors.New("invalid request body")
	ErrExpiredToken   = errors.New("token is expired")
	ErrInvalidToken   = errors.New("token is invalid")
	ErrInvalidKeySize = errors.New("invalid key size")
)
