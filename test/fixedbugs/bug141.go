// run


package main

import "os"

type S struct { i int }
func (p *S) Get() int { return p.i }

type Empty interface {
}

type Getter interface {
	Get() int;
}

func f1(p Empty) {
	switch x := p.(type) {
	default: println("failed to match interface", x); os.Exit(1);
	case Getter: break;
	}

}

func main() {
	var s S;
	f1(&s);
}
