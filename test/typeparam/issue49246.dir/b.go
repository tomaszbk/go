package b

import "./a"

func Crash() { a.Y(a.X)() }
