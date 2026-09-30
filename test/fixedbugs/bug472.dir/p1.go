package p1

import "runtime"

func E() func() int { return runtime.NumCPU }

func F() func() { return runtime.Gosched }

func G() func() string { return runtime.GOROOT }

func H() func() { return runtime.GC }

func I() func() string { return runtime.Version }
