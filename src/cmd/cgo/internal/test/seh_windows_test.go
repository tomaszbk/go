//go:build cgo && windows

package cgotest

import "testing"

func TestCallbackCallersSEH(t *testing.T) { testCallbackCallersSEH(t) }
