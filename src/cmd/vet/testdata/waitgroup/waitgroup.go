// This file contains tests for the waitgroup checker.

package waitgroup

import "sync"

func _() {
	var wg *sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Add(1) // ERROR "WaitGroup.Add called from inside new goroutine"
		defer wg.Done()
		// ...
	}()
	wg.Wait()
}
