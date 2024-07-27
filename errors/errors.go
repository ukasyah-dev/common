package errors

import (
	"net/http"

	"github.com/emitra-labs/common/cases"
)

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	if e.Message == "" {
		return cases.ToSentence(e.Code)
	}
	return e.Message
}

func (e *Error) GetHTTPStatus() int {
	switch e.Code {
	case ALREADY_EXISTS:
		return http.StatusConflict
	case INTERNAL:
		return http.StatusInternalServerError
	case INVALID_ARGUMENT:
		return http.StatusBadRequest
	case NOT_FOUND:
		return http.StatusNotFound
	case PERMISSION_DENIED:
		return http.StatusForbidden
	case UNAUTHENTICATED:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

func pickFirst(msg []string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return ""
}

func AlreadyExists(msg ...string) *Error {
	return &Error{
		Code:    ALREADY_EXISTS,
		Message: pickFirst(msg),
	}
}

func IsAlreadyExists(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == ALREADY_EXISTS
	}
	return false
}

func Internal(msg ...string) *Error {
	return &Error{
		Code:    INTERNAL,
		Message: pickFirst(msg),
	}
}

func IsInternal(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == INTERNAL
	}
	return false
}

func InvalidArgument(msg ...string) *Error {
	return &Error{
		Code:    INVALID_ARGUMENT,
		Message: pickFirst(msg),
	}
}

func IsInvalidArgument(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == INVALID_ARGUMENT
	}
	return false
}

func NotFound(msg ...string) *Error {
	return &Error{
		Code:    NOT_FOUND,
		Message: pickFirst(msg),
	}
}

func IsNotFound(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == NOT_FOUND
	}
	return false
}

func PermissionDenied(msg ...string) *Error {
	return &Error{
		Code:    PERMISSION_DENIED,
		Message: pickFirst(msg),
	}
}

func IsPermissionDenied(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == PERMISSION_DENIED
	}
	return false
}

func Unauthenticated(msg ...string) *Error {
	return &Error{
		Code:    UNAUTHENTICATED,
		Message: pickFirst(msg),
	}
}

func IsUnauthenticated(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == UNAUTHENTICATED
	}
	return false
}
