package f1

type Foo struct {
	doneChan chan bool
	Name     string
	fOO      int
	hook     func()
}

func (f *Foo) Exported() {
}

func (f *Foo) unexported() {
}
