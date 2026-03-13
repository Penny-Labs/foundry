package apperr

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// StatusError is an error that carries an HTTP status and safe client-facing message.
type StatusError interface {
	error
	StatusCode() int
	SafeMessage() string
}

// Error is a status-aware application error.
type Error struct {
	statusCode  int
	safeMessage string
	err         error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.err != nil {
		return e.err.Error()
	}
	return e.safeMessage
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *Error) StatusCode() int {
	if e == nil || e.statusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.statusCode
}

func (e *Error) SafeMessage() string {
	if e == nil || strings.TrimSpace(e.safeMessage) == "" {
		return http.StatusText(http.StatusInternalServerError)
	}
	return e.safeMessage
}

func New(statusCode int, safeMessage string, err error) error {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	if strings.TrimSpace(safeMessage) == "" {
		safeMessage = http.StatusText(statusCode)
	}
	return &Error{
		statusCode:  statusCode,
		safeMessage: safeMessage,
		err:         err,
	}
}

func Invalid(safeMessage string, err error) error {
	return New(http.StatusBadRequest, safeMessage, err)
}

func Unauthorized(safeMessage string, err error) error {
	return New(http.StatusUnauthorized, safeMessage, err)
}

func Forbidden(safeMessage string, err error) error {
	return New(http.StatusForbidden, safeMessage, err)
}

func NotFound(safeMessage string, err error) error {
	return New(http.StatusNotFound, safeMessage, err)
}

func Conflict(safeMessage string, err error) error {
	return New(http.StatusConflict, safeMessage, err)
}

func Internal(err error) error {
	return New(http.StatusInternalServerError, "internal server error", err)
}

// WrapInternal adds operation context while preserving a generic safe internal message.
func WrapInternal(operation string, err error) error {
	if err == nil {
		return nil
	}
	if strings.TrimSpace(operation) == "" {
		return Internal(err)
	}
	return Internal(fmt.Errorf("%s: %w", operation, err))
}

// StatusAndMessage maps a plain error to an HTTP status and safe client message.
func StatusAndMessage(err error) (int, string) {
	var se StatusError
	if errors.As(err, &se) {
		return se.StatusCode(), se.SafeMessage()
	}
	return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
}
