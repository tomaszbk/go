// errorcheck


// Test that we type-check deferred/go functions even
// if they are not called (a common error). Specifically,
// we don't want to see errors such as import or variable
// declared and not used.

package p

import (
	"fmt"
	"math"
)

func f() {
	var i int
	defer func() { fmt.Println() } // ERROR "must be function call"
	go func() { _ = math.Sin(0) }  // ERROR "must be function call"
	go func() { _ = i}             // ERROR "must be function call"
}
