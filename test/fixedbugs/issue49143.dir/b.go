package b

import "./a"


type Loaders struct {
	Loader *a.Loader[int, int]
}

func NewLoaders() *Loaders {
	return new(Loaders)
}
