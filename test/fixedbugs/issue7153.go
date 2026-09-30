// errorcheck

// Issue 7153: array invalid index error duplicated on successive bad values

package p

var _ = []int{a: true, true} // ERROR "undefined: a" "cannot use true \(type untyped bool\) as type int in slice literal|undefined name .*a|incompatible type|cannot use"
