package main

import "featuretest/lib"

func main() {
	v := 7
	if lib.Pick(nil, &v) != &v || lib.Generic((*int)(nil), &v) != &v || lib.Get(nil) != 3 || lib.Get(&lib.Node{Next: &lib.Node{N: 9}}) != 9 {
		panic("export nullsafety")
	}
	m := map[int]*int{}
	lib.Init(m, 0, &v)
	lib.Init(m, 0, nil)
	if m[0] != &v {
		panic("export assignment")
	}
}
