//go:build ignore

package main

import (
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("issue.22295.so")
	if err != nil {
		log.Fatal(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		log.Fatal(err)
	}
	const want = 2503
	got := f.(func() int)()
	if got != want {
		log.Fatalf("got %d, want %d", got, want)
	}
}
