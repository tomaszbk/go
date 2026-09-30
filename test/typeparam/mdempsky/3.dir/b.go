package b

import "./a"

func g() { a.F(make(chan int)) }
