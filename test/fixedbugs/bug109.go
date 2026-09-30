// compile

package bug109

func f(a float64) float64 {
	e := 1.0
	e = e * a
	return e
}

/*
6g bugs/bug109.go
bugs/bug109.go:5: illegal types for operand: MUL
	(<float64>FLOAT64)
	(<float32>FLOAT32)
bugs/bug109.go:5: illegal types for operand: AS
	(<float64>FLOAT64)
bugs/bug109.go:6: illegal types for operand: RETURN
	(<float32>FLOAT32)
	(<float64>FLOAT64)
*/
