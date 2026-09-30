// run


package main

func main() {
	{
		i := I(A{})

		b := make(chan I, 1)
		b <- B{}

		var ok bool
		i, ok = <-b
		_ = ok

		i.M()
	}

	{
		i := I(A{})

		b := make(chan I, 1)
		b <- B{}

		select {
		case i = <-b:
		}

		i.M()
	}

	{
		i := I(A{})

		b := make(chan I, 1)
		b <- B{}

		var ok bool
		select {
		case i, ok = <-b:
		}
		_ = ok

		i.M()
	}
}

type I interface{ M() int }

type T int

func (T) M() int { return 0 }

type A struct{ T }
type B struct{ T }
