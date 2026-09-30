// compile


package foo

func f(x []interface{}) (err error) {
	for _, d := range x {
		_, ok := d.(*int)
		if ok {
			return
		}
	}
	return
}
