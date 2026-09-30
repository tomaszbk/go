// compile


// Issue 45693: ICE with register args.

package p

func f() {
	var s string
	s = s + "" + s + "" + s + ""
	for {
	}
}
