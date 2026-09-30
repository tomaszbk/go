// run

// Test var x = x + 1 works.

package main

func main() {
	var x int = 1
	if x != 1 {
		print("found ", x, ", expected 1\n")
		panic("fail")
	}
	{
		var x int = x + 1
		if x != 2 {
			print("found ", x, ", expected 2\n")
			panic("fail")
		}
	}
	{
		x := x + 1
		if x != 2 {
			print("found ", x, ", expected 2\n")
			panic("fail")
		}
	}
}
