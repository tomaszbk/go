// run


package main


func Alloc(i int) int {
	switch i {
	default:
		return 5;
	case 1:
		return 1;
	case 10:
		return 10;
	}
	return 0
}

func main() {
	s := Alloc(7);
	if s != 5 { panic("bad") }
}

/*
bug028.go:7: unreachable statements in a switch
*/
