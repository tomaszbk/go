package main

import (
	"fmt"
	"reflect"

	"./a"
)

func main() {
	e := []int{1, 2, 2, 3, 1, 6}

	got := a.Unique(e)
	want := []int{1, 2, 3, 6}
	if !reflect.DeepEqual(got, want) {
		panic(fmt.Sprintf("got %d, want %d", got, want))
	}

}
