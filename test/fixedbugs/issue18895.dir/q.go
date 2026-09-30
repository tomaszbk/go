package q

import "./p"

func x() { // ERROR "can inline x"
	p.F() // ERROR "inlining call to .*\.F" "inlining call to .*\.m"
}
