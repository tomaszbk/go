// compile


package issue22198

func f(a *bool, b bool) {
	if b {
		return
	}
	c := '\n'
	if b {
		c = ' '
	}
	*a = c == '\n'
}
