// run

package main

const (
	joao = "João"
	jose = "José"
)

func main() {
	s1 := joao
	s2 := jose
	if (s1 < s2) != (joao < jose) {
		panic("unequal")
	}
}
