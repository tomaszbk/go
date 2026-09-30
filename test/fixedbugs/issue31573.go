// errorcheck -0 -m -l

package p

func f(...*int) {}

func g() {
	defer f()
	defer f(new(int))           // ERROR "... argument does not escape$" "new\(int\) does not escape$"
	defer f(new(int), new(int)) // ERROR "... argument does not escape$" "new\(int\) does not escape$"

	defer f(nil...)
	defer f([]*int{}...)                   // ERROR "\[\]\*int{} does not escape$"
	defer f([]*int{new(int)}...)           // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"
	defer f([]*int{new(int), new(int)}...) // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"

	go f()
	go f(new(int))           // ERROR "... argument does not escape$" "new\(int\) does not escape$"
	go f(new(int), new(int)) // ERROR "... argument does not escape$" "new\(int\) does not escape$"

	go f(nil...)
	go f([]*int{}...)                   // ERROR "\[\]\*int{} does not escape$"
	go f([]*int{new(int)}...)           // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"
	go f([]*int{new(int), new(int)}...) // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"

	for {
		defer f()
		defer f(new(int))           // ERROR "... argument does not escape$" "new\(int\) does not escape$"
		defer f(new(int), new(int)) // ERROR "... argument does not escape$" "new\(int\) does not escape$"

		defer f(nil...)
		defer f([]*int{}...)                   // ERROR "\[\]\*int{} does not escape$"
		defer f([]*int{new(int)}...)           // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"
		defer f([]*int{new(int), new(int)}...) // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"

		go f()
		go f(new(int))           // ERROR "... argument does not escape$" "new\(int\) does not escape$"
		go f(new(int), new(int)) // ERROR "... argument does not escape$" "new\(int\) does not escape$"

		go f(nil...)
		go f([]*int{}...)                   // ERROR "\[\]\*int{} does not escape$"
		go f([]*int{new(int)}...)           // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"
		go f([]*int{new(int), new(int)}...) // ERROR "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"
	}
}
