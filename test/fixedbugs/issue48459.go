// compile


package main

func main() {
	if true {
		return
	}

	defer func() {
		recover()
	}()
}
