// errorcheck

package p

var f = func() { f() } // ERROR "initialization cycle|initialization expression for .*f.* depends upon itself"
