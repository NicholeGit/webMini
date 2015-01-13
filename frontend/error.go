package frontend

import (
	"fmt"
)

const (
	ErrCodeOK                  = 0
	ErrCodeAuthenticateFailed  = 1
	ErrCodeBadRequest          = 400
	ErrCodeUnauthenticated     = 401
	ErrCodeInternalServerError = 500
)

type Error struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func NewError(ErrCode int, ErrMsg string) *Error {
	return &Error{
		ErrCode: ErrCode,
		ErrMsg:  ErrMsg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("err_code: %d, err_msg: %s", e.ErrCode, e.ErrMsg)
}
