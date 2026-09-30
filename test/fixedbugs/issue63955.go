// compile

package j

func f(try func() int, shouldInc func() bool, N func(int) int) {
	var n int
loop: // we want to have 3 preds here, the function entry and both gotos
	if v := try(); v == 42 || v == 1337 { // the two || are to trick findIndVar
		if n < 30 { // this aims to be the matched block
			if shouldInc() {
				n++
				goto loop
			}
			n = N(n) // try to prevent some block joining
			goto loop
		}
	}
}
