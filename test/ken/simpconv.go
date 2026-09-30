// run

// Test simple arithmetic conversion.

package main

type vlong int64
type short int16

func main() {
	s1 := vlong(0)
	for i := short(0); i < 10; i = i + 1 {
		s1 = s1 + vlong(i)
	}
	if s1 != 45 {
		panic(s1)
	}

	s2 := float64(0)
	for i := 0; i < 10; i = i + 1 {
		s2 = s2 + float64(i)
	}
	if s2 != 45 {
		panic(s2)
	}
}
