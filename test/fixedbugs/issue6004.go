// errorcheck

package main

func main() {
	_ = nil                  // ERROR "use of untyped nil"
	_, _ = nil, 1            // ERROR "use of untyped nil"
	_, _ = 1, nil            // ERROR "use of untyped nil"
	_ = append(nil, 1, 2, 3) // ERROR "untyped nil|nil"
}
