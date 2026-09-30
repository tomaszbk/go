package p

func /* ERROR unexpected {, expected name or \($ */ {}
func (T) /* ERROR unexpected {, expected name$ */ {}
func (T) /* ERROR unexpected \(, expected name$ */ () {}
func (T) /* ERROR unexpected \(, expected name$ */ ()
