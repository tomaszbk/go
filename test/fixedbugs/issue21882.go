// errorcheck

package p

type T [2]T // ERROR "invalid recursive type"
