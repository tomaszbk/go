// errorcheck


package main

fmt.Printf("hello")	// ERROR "non-declaration statement outside function body|expected declaration"

func main() {
}

x++	// ERROR "non-declaration statement outside function body|expected declaration"

func init() {
}

x,y := 1, 2	// ERROR "non-declaration statement outside function body|expected declaration"

