package p1

import "fmt"

type Fer interface {
	f() string
}

type Object struct{}

func (this *Object) f() string {
	return "Object.f"
}

func PrintFer(fer Fer) {
	fmt.Sprintln(fer.f())
}
