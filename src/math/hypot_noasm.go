//go:build !386 && !amd64 && !loong64

package math

const haveArchHypot = false

func archHypot(p, q float64) float64 {
	panic("not implemented")
}
