// run


package main

func main() {
	m := "0"
	for _, c := range "321" {
		m = max(string(c), m)
		println(m)
	}
}
