package obj

// go-specific code shared across loaders (5l, 6l, 8l).

func Nopout(p *Prog) {
	p.As = ANOP
	p.Scond = 0
	p.From = Addr{}
	p.RestArgs = nil
	p.Reg = 0
	p.To = Addr{}
}
