package b

import "./a"

type X int

func (X) M() a.T[X] { return 0 }
