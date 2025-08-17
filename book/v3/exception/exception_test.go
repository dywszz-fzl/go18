package exception_test

import (
	"awesomeProject/book/v3/exception"

	"testing"
)

func CheckIsError() error {
	return exception.ErrNotFound("book %d not found", 1)
}

func TestException(t *testing.T) {
	err := CheckIsError()
	t.Log(err)

	if v, ok := err.(*exception.ApiException); ok {
		t.Log(v.Code)
		t.Log(v.Message)
	}

	t.Log(exception.IsApiException(err, exception.CODE_NOT_FOUND))
}
