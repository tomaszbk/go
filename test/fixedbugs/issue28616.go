// compile

// Make sure we don't dead code eliminate a label.

package p

var i int

func f() {

	if true {

		if i == 1 {
			goto label
		}

		return
	}

label:
}
