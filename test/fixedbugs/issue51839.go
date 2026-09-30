// compile


package main

func main() {
	testRecover()

}

func testRecover() {
	if false {
		func() {
			defer func() {
				_ = recover()
			}()
		}()
	}
}
