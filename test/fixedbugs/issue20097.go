// compile


// Issue 20097: ensure that we CSE multiple Select ops with
// the same underlying type

package main

type T int64

func f(x, y int64) (int64, T) {
	a := x / y
	b := T(x) / T(y)
	return a, b
}
