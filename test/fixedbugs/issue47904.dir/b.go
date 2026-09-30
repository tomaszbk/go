package b

import "./a"

// Package b needs [2]a.K with its algorithms. The linker loads a first,
// so without a rule to prefer this descriptor over a's, a's wins and the
// type ends up without an equality function.
var boxed any = [2]a.K{{1, 2}, {3, 4}}

func Get() any { return boxed }
