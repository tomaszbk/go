// compile


package bug118

func Send(c chan int) int {
	select {
	default:
		return 1;
	}
	return 2;
}
