//go:build cgo

package cgotest

import (
	"testing"

	"cmd/cgo/internal/test/issue26213"
)

func test26213(t *testing.T) {
	issue26213.Test26213(t)
}
