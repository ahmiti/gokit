package errors

import (
	"fmt"
)

type Kind int

const (
	NotFound Kind = iota
	Validation
	Conflict
	Unauthorized
	Forbidden
	RateLimited
	Unavailable
	Internal
)

func (k Kind) String() string {
	return [...]string{"NotFound", "Validation", "Conflict", "Unauthorized", "Forbidden", "RateLimited", "Unavailable", "Internal"}[k]
}

type Error struct {
	kind    Kind
	code    string
	message string
	cause   error
}

func New(kind Kind, code, message string) *Error {
	return &Error{
		kind:    kind,
		code:    code,
		message: message,
	}
}

func (e *Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	}
	return e.message
}

func (e *Error) Kind() Kind        { return e.kind }
func (e *Error) Code() string      { return e.code }
func (e *Error) Message() string   { return e.message }
func (e *Error) Cause() error      { return e.cause }

func (e *Error) WithCause(cause error) *Error {
	e.cause = cause
	return e
}

func (e *Error) Unwrap() error {
	return e.cause
}

func HTTPStatus(kind Kind) int {
	switch kind {
	case NotFound:
		return 404
	case Validation:
		return 400
	case Conflict:
		return 409
	case Unauthorized:
		return 401
	case Forbidden:
		return 403
	case RateLimited:
		return 429
	case Unavailable:
		return 503
	default:
		return 500
	}
}
