package b

import "./a"

func F(err error) bool {
	return a.IsMyError(err)
}
