package a

type T struct {
	X int `go:"track"`
	Y int `go:"track"`
	Z int // untracked
}

func (t *T) GetX() int {
	return t.X
}
func (t *T) GetY() int {
	return t.Y
}
func (t *T) GetZ() int {
	return t.Z
}
