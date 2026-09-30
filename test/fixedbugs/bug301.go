// compile

// https://golang.org/issue/990

package main

func main() {
	defer func() {
		if recover() != nil {
			panic("non-nil recover")
		}
	}()
	panic(nil)
}
