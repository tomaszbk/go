// errorcheck


// issue 5172: spurious warn about type conversion on broken type inside go and defer

package main

type foo struct {
	x bar // ERROR "undefined"
}

type T struct{}

func (t T) Bar() {}

func main() {
	var f foo
	go f.bar()    // ERROR "undefined"
	defer f.bar() // ERROR "undefined"

	t := T{1} // ERROR "too many"
	go t.Bar()
}
