package internal

import "errors"

var (
	ErrAbortHandler    = errors.New("net/http: abort Handler")
	ErrBodyNotAllowed  = errors.New("http: request method or response status code does not allow body")
	ErrRequestCanceled = errors.New("net/http: request canceled")
	ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")
)
