package main

import (
	"./a"
	"fmt"
)

func main() {
	_ = a.UnmarshalOptions1{
		Unmarshalers: a.UnmarshalFuncV2(func(opts a.UnmarshalOptions1, dec *a.Decoder1, val *interface{}) (err error) {
			return fmt.Errorf("error")
		}),
	}
}
