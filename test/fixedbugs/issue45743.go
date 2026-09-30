// compile


package main

func fn() func(interface{}) {
	return func(o interface{}) {
		switch v := o.(type) {
		case *int:
			*v = 1
		}
	}
}

func main() {
	fn()
}
