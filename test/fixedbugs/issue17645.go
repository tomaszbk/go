// errorcheck

package main

type Foo struct {
	X int
}

func main() {
	var s []int
	var _ string = append(s, Foo{""}) // ERROR "cannot use append\(s, Foo{…}\) .* as string value in variable declaration" "cannot use Foo{…} .* as int value in argument to append" "cannot use .* as int value in struct literal"
}
