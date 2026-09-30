package b

import "./a"

func F2() int {
	var mia a.MyIntAlias
	return mia.Get()
}
