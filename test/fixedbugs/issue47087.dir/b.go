package b

func F() interface{} { return struct{ _ []int }{} }

var X = F()
