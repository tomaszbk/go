// asmcheck


package codegen

import "unique"

func BytesToHandle(b []byte) unique.Handle[string] {
	// amd64:-`.*runtime\.slicebytetostring\(`
	return unique.Make(string(b))
}

type Pair struct {
	S1 string
	S2 string
}

func BytesPairToHandle(b1, b2 []byte) unique.Handle[Pair] {
	// TODO: should not copy b1 and b2.
	return unique.Make(Pair{string(b1), string(b2)})
}
