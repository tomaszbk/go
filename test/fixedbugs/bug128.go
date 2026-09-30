// run


package main
func main() {
	switch {
		// empty switch is allowed according to syntax
		// unclear why it shouldn't be allowed
	}
	switch tag := 0; tag {
		// empty switch is allowed according to syntax
		// unclear why it shouldn't be allowed
	}
}

/*
uetli:~/Source/go1/test/bugs gri$ 6g bug127.go 
bug127.go:5: switch statement must have case labels
bug127.go:9: switch statement must have case labels
*/
