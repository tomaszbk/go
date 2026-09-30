package main

import "./p"

func main() {
	_ = p.MarshalFuncV1[int](func(int) ([]byte, error) { return nil, nil })
}
