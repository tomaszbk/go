// errorcheck


package main
func main() {
	var v interface{} = 0;
	switch v.(type) {
	case int:
		fallthrough;		// ERROR "fallthrough"
	default:
		panic("fell through");
	}
}
