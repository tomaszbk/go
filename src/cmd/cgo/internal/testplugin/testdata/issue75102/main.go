package main

import (
	"fmt"
	"plugin"
)

func init() {
	_, err := plugin.Open("issue75102.so")
	if err == nil {
		panic("unexpected success to open a different version plugin")
	}
}

func main() {
	fmt.Println("done")
}
