// errorcheck


package main

func main() {
	_ = string(-4 + 2i + 2) // ERROR "-4 \+ 2i|invalid type conversion"
}
