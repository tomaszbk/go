// run


package main

func f[G any]() int {
	return func() int {
		return func() int {
			return 0
		}()
	}()
}

func main() {
	f[int]()
}
