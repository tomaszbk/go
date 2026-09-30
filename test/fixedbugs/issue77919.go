// compile


package main

var sink any

func main() {
	i := 0
	output := make([]string, 8, i)
	sink = output
}
