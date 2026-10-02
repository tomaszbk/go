package main

import (
	"conditional/lib"
	"fmt"
	"reflect"
	"strconv"
)

func check(label string, got, want any) {
	if !reflect.DeepEqual(got, want) {
		panic(fmt.Sprintf("%s: got %#v, want %#v", label, got, want))
	}
	fmt.Printf("%s: %v\n", label, got)
}

// The exported constant can size an array.
var buffer [lib.Size / 8]byte

func main() {
	for _, c := range []bool{true, false} {
		calls := 0
		then := func() int { calls++; return 10 }
		els := func() int { calls += 100; return 20 }
		check("pick", lib.Pick(c, 1, 2), map[bool]int{true: 1, false: 2}[c])
		check("label", lib.Label(map[bool]int{true: 1, false: 2}[c]), map[bool]string{true: "item", false: "items"}[c])
		check("lazy", []int{lib.Lazy(c, then, els), calls}, map[bool][]int{true: {10, 1}, false: {20, 100}}[c])
		check("typed nil", lib.ErrorOf(c) == nil, !c)
		for _, fail := range []bool{false, true} {
			v, err := lib.Value(c, fail)
			want := []any{8, nil}
			switch {
			case !c:
				want = []any{0, nil}
			case fail:
				want = []any{0, lib.Failure}
			}
			check("value", []any{v, err}, want)
		}
		n, zero := 5, 0
		check("generic", []any{lib.Generic(c, "a", "b"), lib.Generic(c, 1.5, 2.5), *lib.Generic(c, &n, &zero)},
			map[bool][]any{true: {"a", 1.5, 5}, false: {"b", 2.5, 0}}[c])
		key := map[bool]string{true: "k", false: "missing"}[c]
		check("or default", lib.OrDefault(map[string][]int{"k": {1}}, key, nil) == nil, !c)
		check("generic method", []any{lib.Box[string]{V: "v", OK: c}.Get("def"), lib.Box[int]{V: 1, OK: c}.Get(-1)},
			map[bool][]any{true: {"v", 1}, false: {"def", -1}}[c])
	}
	check("constants", []any{lib.Size, len(buffer), lib.Mode},
		[]any{strconv.IntSize, strconv.IntSize / 8, map[bool]string{true: "wide", false: "narrow"}[strconv.IntSize == 64]})
}
