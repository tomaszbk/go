// errorcheck

// Issue 1871.

package p

type a interface {
	foo(x int) (x int) // ERROR "duplicate argument|redefinition|redeclared"
}

/*
Previously:

bug.go:1 x redclared in this block
    previous declaration at bug.go:1
*/
