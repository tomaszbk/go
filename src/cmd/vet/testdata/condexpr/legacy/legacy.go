// The scenarios of ../modern/modern.go written with if statements, as with
// standard Go. Each function must get the same diagnostics as its modern
// counterpart (see TestCondExpr).

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
	if c {
		return use(cancel)
	}
	return 0 // ERROR `this return statement may be reached without using the cancel var defined on line \d+`
}

func usedOnBothBranches(c bool) context.CancelFunc {
	_, cancel := context.WithCancel(background())
	if c {
		return cancel
	}
	return func() { cancel() }
}

func usedInCondition() int {
	_, cancel := context.WithCancel(background())
	if use(cancel) > 0 {
		return 1
	}
	return 2
}

func handlerInBranch(c bool) (int, error) {
	_, cancel := context.WithCancel(background()) // ERROR `the cancel function is not used on all paths \(possible context leak\)`
	var n int
	if c {
		v, err := source()
		if err != nil {
			return 0, err // ERROR `this return statement may be reached without using the cancel var defined on line \d+`
		}
		n = v
	} else {
		n = 0
	}
	cancel()
	return n, nil
}

func handlerInLiteralInBranch(c bool) func() error {
	_, cancel := context.WithCancel(background())
	if c {
		return func() error {
			if err := check(); err != nil {
				cancel()
				return err
			}
			return nil
		}
	}
	return func() error { cancel(); return nil }
}

// copylocks: only the selected branch is copied.

func copyFromBranch(c bool, a *sync.Mutex) *sync.Mutex {
	var x sync.Mutex
	if c {
		x = *a // ERROR "assignment copies lock value to x: sync.Mutex"
	} else {
		x = sync.Mutex{}
	}
	return &x
}

func noCopy(c bool) *sync.Mutex {
	var x sync.Mutex
	if c {
		x = sync.Mutex{}
	} else {
		x = newMutex()
	}
	return &x
}

func interfaceTarget(c bool, a *sync.Mutex) any {
	var z any
	if c {
		z = *a // ERROR "assignment copies lock value to z: sync.Mutex"
	} else {
		z = nil
	}
	return z
}

func callArgument(c bool, a *sync.Mutex) {
	if c {
		consume(*a) // ERROR "call of consume copies lock value: sync.Mutex"
	} else {
		consume(sync.Mutex{})
	}
}

func resultOfCall(c bool) (*sync.Mutex, error) {
	var w sync.Mutex
	if c {
		var err error
		w, err = mutexOrErr()
		if err != nil {
			return nil, err
		}
	} else {
		w = sync.Mutex{}
	}
	return &w, nil
}

// shift: a constant condition makes the other branch dead code.

func deadBranch(x uint32) uint32 {
	const debug = false
	var y uint32
	if debug {
		y = x << 40
	} else {
		y = x
	}
	return y
}

func liveBranch(x uint32) uint32 {
	const on = true
	var y uint32
	if on {
		y = x << 40 // ERROR "x .32 bits. too small for shift of 40"
	} else {
		y = x
	}
	return y
}
