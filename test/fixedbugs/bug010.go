// run


package main


func f(i int, f float64) {
	i = 8
	f = 8.0
	return
}

func main() {
	f(3, float64(5))
}

/*
bug10.go:5: i undefined
bug10.go:6: illegal conversion of constant to 020({},<_o001>{<i><int32>INT32;<f><float32>FLOAT32;},{})
bug10.go:7: error in shape across assignment
*/
