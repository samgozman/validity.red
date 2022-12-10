package main

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

var (
	ErrUnauthorized     = errors.New("authentication failed")
	ErrInvalidInputs    = errors.New("invalid inputs")
	ErrInvalidCaptcha   = errors.New("invalid captcha")
	ErrEmailNotVerified = errors.New("email not verified")
)

var ErrorsArr = []error{
	ErrUnauthorized,
	ErrInvalidInputs,
	ErrInvalidCaptcha,
	ErrEmailNotVerified,
}

// Map error types to HTTP status codes
var ErrorStatus = map[error]int{
	ErrUnauthorized:     http.StatusUnauthorized,
	ErrInvalidInputs:    http.StatusBadRequest,
	ErrInvalidCaptcha:   http.StatusBadRequest,
	ErrEmailNotVerified: http.StatusUnauthorized,
}

// Map gRPC codes to HTTP status codes.
var RPCStatus = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Canceled:           http.StatusConflict, // ?
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusRequestTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusRequestedRangeNotSatisfiable,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusUnprocessableEntity,
	codes.Unauthenticated:    http.StatusUnauthorized,
}
