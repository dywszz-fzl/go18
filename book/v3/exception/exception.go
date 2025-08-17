package exception

import (
	"errors"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

// 用于描述业务异常
// 实现自定义异常
type ApiException struct {
	Code int `json:"code"`

	Message string `json:"message"`

	HttpCode int `json:"-"`
}

func (e *ApiException) Error() string {
	return e.Message
}

func (e *ApiException) String() string {
	return pretty.ToJSON(e)
}

func (e *ApiException) WithMessage(msg string) *ApiException {
	e.Message = msg
	return e
}

func (e *ApiException) WithHttpCode(code int) *ApiException {
	e.HttpCode = code
	return e
}
func IsApiException(err error, code int) bool {
	var apiErr *ApiException
	if errors.As(err, &apiErr) {
		return apiErr.Code == code
	}
	return false
}
