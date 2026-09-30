package main

type N[B N[B]] interface {
	Add(B) B
}

func Add[P N[P]](x, y P) P {
	return x.Add(y)
}

type MyInt int

func (x MyInt) Add(y MyInt) MyInt {
	return x + y
}

func main() {
	var x, y MyInt = 2, 3
	println(Add(x, y))
}
