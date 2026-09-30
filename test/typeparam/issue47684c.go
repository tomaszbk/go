// run


package main

func f[G any]() func()func()int {
	return func() func()int {
		return func() int {
			return 0
		}
	}
}

func main() {
	f[int]()()()
}
