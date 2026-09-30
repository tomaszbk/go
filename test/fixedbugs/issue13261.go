// compile


// Taking the address of a parenthesized composite literal is permitted.

package main

type T struct{}

func main() {
	_ = &T{}
	_ = &(T{})
	_ = &((T{}))

	_ = &struct{}{}
	_ = &(struct{}{})
	_ = &((struct{}{}))

	switch (&T{}) {}
	switch &(T{}) {}
	switch &((T{})) {}

	switch &struct{}{} {}
	switch &(struct{}{}) {}
	switch &((struct{}{})) {}
}
