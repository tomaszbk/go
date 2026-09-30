// compile

// Issue 11750: mkdotargslice: typecheck failed

package main

func main() {
	fn := func(names string) {

	}
	func(names ...string) {
		for _, name := range names {
			fn(name)
		}
	}("one", "two")
}
