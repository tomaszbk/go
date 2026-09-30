package b

import "reflect"

func B() {
	t1 := reflect.TypeOf([30]int{})
	t2 := reflect.TypeOf(new([30]int)).Elem()
	if t1 != t2 {
		panic("[30]int types do not match")
	}
}
