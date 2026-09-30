package q

import "./p"

func H() {
	p.F() // ERROR "inlining call to p.F"
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
