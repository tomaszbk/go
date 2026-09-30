// Package internal contains definitions used by both log and log/slog.
package internal

// DefaultOutput holds a function which calls the default log.Logger's
// output function.
// It allows slog.defaultHandler to call into an unexported function of
// the log package.
var DefaultOutput func(pc uintptr, data []byte) error
