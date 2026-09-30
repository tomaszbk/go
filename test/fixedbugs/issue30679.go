// compile

package main

func main() {
	var f float64
	var p, q *float64

	p = &f
	if *q > 0 {
		p = q
	}
	_ = *p
}
