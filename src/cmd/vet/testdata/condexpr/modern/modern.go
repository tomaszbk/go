// The scenarios of ../legacy/legacy.go written with conditional expressions.
// Each function must get the same diagnostics as its legacy counterpart
// (see TestCondExpr).

package condexpr

import (
	"context"
	"sync"
)

func source() (int, error)            { return 1, nil }
func check() error                    { return nil }
func use(f context.CancelFunc) int    { f(); return 1 }
func newMutex() sync.Mutex            { return sync.Mutex{} }
func mutexOrErr() (sync.Mutex, error) { return sync.Mutex{}, nil }
func consume(any)                     {}
func background() context.Context     { return context.Background() }

// lostcancel: the branches are alternatives.

func leakOnOneBranch(c bool) int {
	_, cancel := context.WithCancel(background()) // ERROR `the cancel function is not used on all paths \(possible context leak\)`
	return if c { use(cancel) } else { 0 }        // ERROR `this return statement may be reached without using the cancel var defined on line \d+`
}

func usedOnBothBranches(c bool) context.CancelFunc {
	_, cancel := context.WithCancel(background())
	return if c { cancel } else { func() { cancel() } }
}

func usedInCondition() int {
	_, cancel := context.WithCancel(background())
	return if use(cancel) > 0 { 1 } else { 2 }
}

func handlerInBranch(c bool) (int, error) {
	_, cancel := context.WithCancel(background()) // ERROR `the cancel function is not used on all paths \(possible context leak\)`
	n := if c {
		source() or err {
			return 0, err // ERROR `this return statement may be reached without using the cancel var defined on line \d+`
		}
	} else {
		0
	}
	cancel()
	return n, nil
}

func handlerInLiteralInBranch(c bool) func() error {
	_, cancel := context.WithCancel(background())
	return if c {
		func() error {
			check() or err {
				cancel()
				return err
			}
			return nil
		}
	} else {
		func() error { cancel(); return nil }
	}
}

// copylocks: only the selected branch is copied.

func copyFromBranch(c bool, a *sync.Mutex) *sync.Mutex {
	x := if c { *a } else { sync.Mutex{} } // ERROR "assignment copies lock value to x: sync.Mutex"
	return &x
}

func noCopy(c bool) *sync.Mutex {
	x := if c { sync.Mutex{} } else { newMutex() }
	return &x
}

func interfaceTarget(c bool, a *sync.Mutex) any {
	var z any
	z = if c { *a } else { nil } // ERROR "assignment copies lock value to z: sync.Mutex"
	return z
}

func callArgument(c bool, a *sync.Mutex) {
	consume(if c { *a } else { sync.Mutex{} }) // ERROR "call of consume copies lock value: sync.Mutex"
}

func resultOfCall(c bool) (*sync.Mutex, error) {
	w := if c { mutexOrErr()! } else { sync.Mutex{} }
	return &w, nil
}

// shift: a constant condition makes the other branch dead code.

func deadBranch(x uint32) uint32 {
	const debug = false
	return if debug { x << 40 } else { x }
}

func liveBranch(x uint32) uint32 {
	const on = true
	return if on { x << 40 } else { x } // ERROR "x .32 bits. too small for shift of 40"
}
