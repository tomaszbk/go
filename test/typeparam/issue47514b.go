// run


package main

func Do[T any](do func() (T, string)) {
	_ = func() (T, string) {
		return do()
	}
}

func main() {
	Do[int](func() (int, string) {
		return 3, "3"
	})
}
