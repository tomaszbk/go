// run


package main

func main() {	
	c := make(chan bool, 1);
	select {
	case _ = <-c:
		panic("BUG: recv should not");
	default:
	}
	c <- true;
	select {
	case _ = <-c:
	default:
		panic("BUG: recv should");
	}
}
