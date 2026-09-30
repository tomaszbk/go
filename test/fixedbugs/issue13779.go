// errorcheck


// Issue 13779: provide better error message when directly assigning to struct field in map

package main

func main() {
	type person struct{ age, weight, height int }
	students := map[string]person{"sally": person{12, 50, 32}}
	students["sally"].age = 3 // ERROR "cannot assign to struct field .* in map"
}
