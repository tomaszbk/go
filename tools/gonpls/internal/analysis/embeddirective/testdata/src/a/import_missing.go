package a

import (
	"fmt"
)

//go:embed embedtext // want "must import \"embed\" when using go:embed directives"
var s string

// This is main function
func main() {
	fmt.Println(s)
}
