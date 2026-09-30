// compile


// /tmp/x.go:5: cannot use _ as value

package p

func f(ch chan int) bool {
	select {
	case _, ok := <-ch:
		return ok
	}
	_, ok := <-ch
	_ = ok
	select {
	case _, _ = <-ch:
		return true
	}
	return false
}
