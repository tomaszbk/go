package nointerface

type I int

//go:nointerface
func (p *I) Get() int { return int(*p) }

func (p *I) Set(v int) { *p = I(v) }
