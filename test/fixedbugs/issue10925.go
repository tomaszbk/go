// run


package main

import "fmt"

func prototype(xyz []string) {}
func main() {
	var got [][]string
	f := prototype
	f = func(ss []string) { got = append(got, ss) }
	for _, s := range []string{"one", "two", "three"} {
		f([]string{s})
	}
	if got[0][0] != "one" || got[1][0] != "two" || got[2][0] != "three" {
		// Bug's wrong output was [[three] [three] [three]]
		fmt.Println("Expected [[one] [two] [three]], got", got)
	}
}
