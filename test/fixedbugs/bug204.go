// run


package main

func main() {
	nchar := 0
	a := []rune{'日', '本', '語', 0xFFFD}
	for _, char := range "日本語\xc0" {
		if nchar >= len(a) {
			println("BUG")
			break
		}
		if char != a[nchar] {
			println("expected", a[nchar], "got", char)
			println("BUG")
			break
		}
		nchar++
	}
}
