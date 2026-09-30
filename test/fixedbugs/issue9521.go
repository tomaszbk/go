// errorcheck


// Test that an incorrect use of the blank identifier is caught.
// Does not compile.

package main

func f() (_, _ []int)         { return }
func g() (x []int, y float64) { return }

func main() {
	_ = append(f()) // ERROR "cannot use \[\]int value as type int in append|cannot use.*type \[\]int.*to append"
	_ = append(g()) // ERROR "cannot use float64 value as type int in append|cannot use.*type float64.*to append"
}
