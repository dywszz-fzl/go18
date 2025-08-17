package exception

import (
	"fmt"
)

const (
	CODE_SERVER_ERROR  = 5000
	CODE_NOT_FOUND     = 404
	CODE_PARAM_INVALID = 400
)

func ErrNotFound(Format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_NOT_FOUND,
		Message:  fmt.Sprintf(Format, a...),
		HttpCode: 500,
	}
}
func ErrServerInternal(Format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_SERVER_ERROR,
		Message:  fmt.Sprintf(Format, a...),
		HttpCode: 404,
	}
}

func ErrValidateFailed(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_PARAM_INVALID,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: 400,
	}
}
