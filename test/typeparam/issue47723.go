// run


package main

func f[_ any]() int {
	var a [1]int
	_ = func() int {
		return func() int {
			return 0
		}()
	}()
	return a[func() int {
		return 0
	}()]
}

func main() {
	f[int]()
}
