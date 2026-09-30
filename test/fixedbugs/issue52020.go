// compile

package main

func main() {
	var _ interface{} = struct{ _ [1]int8 }{}
}
