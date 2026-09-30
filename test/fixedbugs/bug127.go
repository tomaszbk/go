// errorcheck


package main
func main() {
	var x int64 = 0;
	println(x != nil);	// ERROR "illegal|incompatible|nil"
	println(0 != nil);	// ERROR "illegal|incompatible|nil"
}
