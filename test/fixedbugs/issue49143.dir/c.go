package c

import "./b"

type Resolver struct{}

type todoResolver struct{ *Resolver }

func (r *todoResolver) F() {
	b.NewLoaders().Loader.Load()
}
