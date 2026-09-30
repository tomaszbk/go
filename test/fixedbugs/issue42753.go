// compile -d=ssa/check/on

package main

func f() uint32 {
	s := "\x01"
	x := -int32(s[0])
	return uint32(x) & 0x7fffffff
}
