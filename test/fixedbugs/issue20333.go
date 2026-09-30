// compile


// Issue 20333: early checkwidth of [...] arrays led to compilation errors.

package main

import "fmt"

func main() {
	fmt.Println(&[...]string{"abc", "def", "ghi"})
}
