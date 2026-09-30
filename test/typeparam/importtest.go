// compile

// This file checks that basic importing works in -G mode.

package p

import "fmt"
import "math"

func f(x float64) {
	fmt.Println(math.Sin(x))
}
