// run

// Check static composite literal reports wrong for struct
// field.

package main

type one struct {
	i interface{}
}

type two struct {
	i interface{}
	s []string
}

func main() {
	o := one{i: two{i: 42}.i}
	println(o.i.(int))
}
