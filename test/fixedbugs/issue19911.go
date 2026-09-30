// run

package main

import (
	"fmt"
	"strings"
)

type ET struct{}

func (*ET) Error() string { return "err" }

func main() {
	check("false", fmt.Sprintf("(*ET)(nil) == error(nil): %v", (*ET)(nil) == error(nil)))
	check("true", fmt.Sprintf("(*ET)(nil) != error(nil): %v", (*ET)(nil) != error(nil)))

	nilET := (*ET)(nil)
	nilError := error(nil)

	check("false", fmt.Sprintf("nilET == nilError: %v", nilET == nilError))
	check("true", fmt.Sprintf("nilET != nilError: %v", nilET != nilError))
}

func check(want, gotfull string) {
	got := gotfull[strings.Index(gotfull, ": ")+len(": "):]
	if got != want {
		panic("want " + want + " got " + got + " from " + gotfull)
	}
}
