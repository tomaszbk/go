// run


package main

import (
	"bytes"
	"runtime"
	"runtime/pprof"
	"sync"
)

func test() {
	var wg sync.WaitGroup
	wg.Add(2)
	test := func() {
		for i := 0; i < 10; i++ {
			buf := &bytes.Buffer{}
			pprof.Lookup("goroutine").WriteTo(buf, 2)
		}
		wg.Done()
	}

	go test()
	go test()
	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 10; i++ {
		test()
	}
}
