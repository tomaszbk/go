// errorcheck


package main

func main() {
	s := uint(10)
	ss := 1 << s
	y1 := float64(ss)
	y2 := float64(1 << s) // ERROR "shift"
	y3 := string(1 << s)  // ERROR "argument must have type byte or rune"
	_, _, _, _, _ = s, ss, y1, y2, y3
}
