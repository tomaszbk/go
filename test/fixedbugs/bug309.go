// compile

// issue 1016

package bug309

func foo(t interface{}, c chan int) {
	switch v := t.(type) {
	case int:
		select {
		case <-c:
			// bug was: internal compiler error: var without type, init: v
		}
	default:
		_ = v
	}
}
