// compile

package p

func e() {
	for true {
		if true {
			continue
		}
	}
}

func g() {}

func f() {
	i := 0
	if true {
		i++
	}
	for true {
		continue
		g()
	}
}
