package main

var f *int

func init() {
	f = new(int)
	*f = 2503
}

func F() int { return *f }

func main() {}
