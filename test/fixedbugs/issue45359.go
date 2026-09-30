// compile


package main

func f() {
	var i, j int
	var b bool
	i = -(i &^ i)
	for 1>>uint(i) == 0 {
		_ = func() {
			i, b = 0, true
		}
		_ = b
		i %= j
	}
}
