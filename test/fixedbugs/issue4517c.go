// errorcheck

package p

type init byte // ERROR "cannot declare init - must be func"
