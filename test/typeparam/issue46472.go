// run


package main

func foo[T any](d T) {
	switch v := interface{}(d).(type) {
	case string:
		if v != "x" {
			panic("unexpected v: " + v)
		}
	}

}
func main() {
	foo("x")
}
