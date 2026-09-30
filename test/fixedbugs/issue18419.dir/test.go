// errorcheck -0 -m -l


package main

import "./other"

func InMyCode(e *other.Exported) {
	e.member() // ERROR "e\.member undefined .cannot refer to unexported field or method other\.\(\*Exported\)\.member.|unexported field or method"
}

func main() {}
