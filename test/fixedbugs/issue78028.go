// compile

package main

var z = map[int]int{0: 1}

func main() {
	_ = make([]byte, z[0], 1)
}
