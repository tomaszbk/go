package a

func F() interface{} { return struct{ _ []int }{} }

var X = F()
