// errorcheck


//go:build linux && 386

package p

var _ struct {
	a, b [1 << 30]byte // ERROR "type struct .* too large"
}
