// compile

// PR61244: Type descriptors expressions were not traversed, causing an ICE
// in gccgo when producing the backend representation.
// This is a reduction of a program reported by GoSmith.

package main

const a = 0

func main() {
	switch i := (interface{})(a); i.(type) {
	case [0]string:
	}
}
