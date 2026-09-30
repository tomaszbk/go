//go:build cgo && darwin

package cgotest

import "testing"

func TestIssue76023(t *testing.T)          { issue76023(t) }
func TestGlobalDataDynimport(t *testing.T) { unsignedRelocDynimport(t) }
