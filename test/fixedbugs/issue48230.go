// errorcheck

package p

//go:embed issue48230.go // ERROR `go:embed only allowed in Go files that import "embed"`
var _ string
