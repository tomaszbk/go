package gonanalysis

func read(v int) (int, error)   { return v, nil }
func flag(v bool) (bool, error) { return v, nil }
func flush() error              { return nil }

var choice = if true { 1 } else { 2 }

func Modern(c bool, xs []int, m map[int][]int) (int, error) {
	x := if c {
		read(1)!
	} else {
		read(2) or err {
			return 0, err
		}
	}
	flush() or err {
		_ = err
	}
	flush()!
	f := func() (int, error) { return read(x)! + 1, nil }
	if if c { true } else { false } {
	}
	if flag(c)! {
	}
	if flag(c)! {
	}
	for i := range xs {
		xs[i] = xs[if c { i } else { 0 }]
	}
	if m[if c { 1 } else { 2 }] == nil {
		m[if c { 1 } else { 2 }] = []int{x}
	}
	if c || flag(c)! {
		x++
	}
	_ = func() bool {
		for if c { false } else { false } {
		}
		return c
	}
	return f()!, nil
}
