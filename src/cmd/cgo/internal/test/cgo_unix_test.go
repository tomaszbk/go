//go:build cgo && !windows

package cgotest

import "testing"

func TestSigaltstack(t *testing.T) { testSigaltstack(t) }
func TestSigprocmask(t *testing.T) { testSigprocmask(t) }
func Test18146(t *testing.T)       { test18146(t) }
