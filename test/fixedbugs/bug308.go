// compile

// issue 1136

package main

import "fmt"

func log1(f string, argv ...interface{}) {
	fmt.Printf("log: %s\n", fmt.Sprintf(f, argv...))
}

func main() {
	log1("%d", 42)
}
