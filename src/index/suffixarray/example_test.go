package suffixarray_test

import (
	"fmt"
	"index/suffixarray"
)

func ExampleIndex_Lookup() {
	index := suffixarray.New([]byte("banana"))
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
	}

	// Unordered output:
	// 1
	// 3
}
