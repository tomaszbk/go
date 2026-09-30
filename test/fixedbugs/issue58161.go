// compile -d=ssa/check/seed=1


package main

func F[G int]() int {
	return len(make([]int, copy([]G{}, []G{})))
}

func main() {
	F[int]()
}
