// compile

package main

var x [3][3]int

func main() {
	for i := range 3 {
		x[i][i] = 0
	}
}
