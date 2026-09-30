package a

var S struct {
	Str string `tag`
}

func F() string {
	v := S
	return v.Str
}
