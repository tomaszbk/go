// run


package main

func f(j int) {
loop:
	switch j {
	case 1:
		break loop
	default:
		println(j)
	}
}

func main() {
loop:
	for j := 0; j < 5; j++ {
		f(j)
		if j == 3 {
			break loop
		}
	}
}
