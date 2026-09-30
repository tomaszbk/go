// run

// Check that calling a nil func causes a proper panic.

package main

func main() {
	defer func() {
		err := recover()
		if err == nil {
			panic("panic expected")
		}
	}()

	var f func()
	f()
}
