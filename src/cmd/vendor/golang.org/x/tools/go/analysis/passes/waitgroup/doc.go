// Package waitgroup defines an Analyzer that detects simple misuses
// of sync.WaitGroup.
//
// # Analyzer waitgroup
//
// waitgroup: check for misuses of sync.WaitGroup
//
// This analyzer detects mistaken calls to the (*sync.WaitGroup).Add
// method from inside a new goroutine, causing Add to race with Wait:
//
//	// WRONG
//	var wg sync.WaitGroup
//	go func() {
//	        wg.Add(1) // "WaitGroup.Add called from inside new goroutine"
//	        defer wg.Done()
//	        ...
//	}()
//	wg.Wait() // (may return prematurely before new goroutine starts)
//
// The correct code calls Add before starting the goroutine:
//
//	// RIGHT
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		...
//	}()
//	wg.Wait()
package waitgroup
