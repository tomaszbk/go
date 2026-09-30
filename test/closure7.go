// run


package main

func g(f func()) {
}

// Must have exportable name
func F() {
	g(func() {
		ch := make(chan int)
		for {
			select {
			case <-ch:
				return
			default:
			}
		}
	})
}

func main() {
	F()
}
