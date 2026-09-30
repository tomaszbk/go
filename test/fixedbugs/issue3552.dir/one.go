package one

// Issue 3552

type T struct { int }

func (t T) F() int { return t.int }

type U struct { int int }

func (u U) F() int { return u.int }

type lint int

type V struct { lint }

func (v V) F() int { return int(v.lint) }

type W struct { lint lint }

func (w W) F() int { return int(w.lint) }



