// compile


package p

type T1 struct{}
type T2 struct{}

func f() {
	switch (T1{}) {
	case T1(T2{}):
	}
}
