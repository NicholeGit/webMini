package model

import "errors"

const (
//	ErrCodeOK                  = 0
//	ErrCodeAuthenticateFailed  = 1
//	ErrCodeBadRequest          = 400
//	ErrCodeUnauthenticated     = 401
//	ErrCodeInternalServerError = 500

)

var (
	ErrNotFound = errors.New("item not found")
)
