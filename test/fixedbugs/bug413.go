// errorcheck


package p

func f(i int) int { return i }

var i = func() int {a := f(i); return a}()  // ERROR "initialization cycle|depends upon itself"
