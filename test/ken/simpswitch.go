// run


// Test simple switch.

package main

func main() {
	r := ""
	a := 3
	for i := 0; i < 10; i = i + 1 {
		switch i {
		case 5:
			r += "five"
		case a, 7:
			r += "a"
		default:
			r += string(rune(i) + '0')
		}
		r += "out" + string(rune(i)+'0')
	}
	if r != "0out01out12out2aout34out4fiveout56out6aout78out89out9" {
		panic(r)
	}
}
