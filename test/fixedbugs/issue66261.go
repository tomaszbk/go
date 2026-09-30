// run

package main

func main() {
	env := func() func(*bool) func() int {
		return func() func(*bool) func() int {
			return func(ptr *bool) func() int {
				return func() int {
					*ptr = true
					return 0
				}
			}
		}()
	}()

	var ok bool
	func(int) {}(env(&ok)())
	if !ok {
		panic("FAIL")
	}
}
