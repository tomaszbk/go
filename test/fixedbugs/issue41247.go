// errorcheck

package p

func f() [2]int {
	return [...]int{2: 0} // ERROR "cannot use \[\.\.\.\]int{.*} \(.*type \[3\]int\)"
}
