// compile


// Crashes 6g, 8g
// https://golang.org/issue/238

package main

func main() {
	bar := make(chan bool);
	select {
	case _ = <-bar:
		return
	}
}

/*
6g bug218.go 
<epoch>: fatal error: dowidth: unknown type: blank
*/
