// compile


package main

type Foo interface { }

type T struct {}
func (t *T) foo() {}

func main() {
	t := new(T);
	var i interface {};
	f, ok := i.(Foo);
	_, _, _ = t, f, ok;
}
