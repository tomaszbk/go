// compile

// Issue 28390/28430: Function call arguments were not
// converted correctly under some circumstances.

package main

func g(_ interface{}, e error)
func h() (int, error)

func f() {
	g(h())
}
