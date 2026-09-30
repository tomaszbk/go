package recursiveiter

import "iter"

type cons struct {
	car int
	cdr *cons
}

func (cons *cons) All() iter.Seq[int] {
	return func(yield func(int) bool) {
		// The correct recursion is:
		//   func (cons *cons) all(f func(int) bool) {
		//     return cons == nil || yield(cons.car) && cons.cdr.all()
		//   }
		// then:
		//   _ = cons.all(yield)
		if cons != nil && yield(cons.car) {
			for elem := range cons.All() { // want "inefficient recursion in iterator All"
				if !yield(elem) {
					break
				}
			}
		}
	}
}
