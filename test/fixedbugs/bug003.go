// run


package main

func main() {
	switch ; { case true: return; default: return }
}
/*
bug003.go:6: fatal error: walkswitch: not case EMPTY
*/
