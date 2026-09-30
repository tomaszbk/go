package main

import (
	"runtime"
	"sync"
)

func init() {
	register("PanicRace", PanicRace)
}

func PanicRace() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			runtime.Gosched()
		}()
		panic("crash")
	}()
	wg.Wait()
}
