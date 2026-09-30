// This file contains tests for the slog checker.

package slog

import "log/slog"

func SlogTest() {
	slog.Info("msg", "a") // ERROR "call to slog.Info missing a final value"
}
