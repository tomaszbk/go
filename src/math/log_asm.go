//go:build amd64 || loong64 || s390x

package math

const haveArchLog = true

func archLog(x float64) float64
