// errorcheck

package main

func foo() {
	_ = func() {}
}

func foo() { // ERROR "foo redeclared in this block|redefinition of .*foo.*"
	_ = func() {}
}

func main() {}
