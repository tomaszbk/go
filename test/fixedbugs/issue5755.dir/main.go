package main

import "./a"

func main() {
	a.Test1("frumious")
	a.Test2("frumious")
	a.Test3("frumious")
	a.Test4("frumious")

	a.Test5(nil)
	a.Test6(nil)
	a.Test7(nil)
	a.Test8(nil)
	a.Test9(0)

	a.TestBar()
	a.IsBaz(nil)
}
