// run

package main

func main() {
	d := diff([]int{}, func(int) string {
		return "foo"
	})
	d()
}

func diff[T any](previous []T, uniqueKey func(T) string) func() {
	return func() {
		newJSON := map[string]T{}
		for _, prev := range previous {
			delete(newJSON, uniqueKey(prev))
		}
	}
}
