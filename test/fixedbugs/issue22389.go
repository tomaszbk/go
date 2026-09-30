// errorcheck -d=panic

package p

type Foo struct{}

func (f *Foo) Call(cb func(*Foo)) {
	cb(f)
}

func main() {
	f := &Foo{}
	f.Call(func(f) {}) // ERROR "f .*is not a type"
}
