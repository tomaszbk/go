//go:build 386 || amd64 || loong64

package math

const haveArchHypot = true

func archHypot(p, q float64) float64
