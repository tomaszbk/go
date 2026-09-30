package errorsas

import "errors"

type MyError struct{}

func (MyError) Error() string { return "" }

func _(err error) {
	var target MyError
	errors.As(err, target) // ERROR "second argument to errors.As must be a non-nil pointer"
}
