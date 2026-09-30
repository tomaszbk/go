// run


package main

func main() {
	x := 0
	func() {
		x = 1
	}()
	func() {
		if x != 1 {
			panic("x != 1")
		}
	}()
}