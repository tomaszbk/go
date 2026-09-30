package a

type T struct { x int }

func F() interface{} {
	return [2]T{}
}

func P() interface{} {
	return &[2]T{}
}
