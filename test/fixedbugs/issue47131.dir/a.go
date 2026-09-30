package a

type MyInt int

type MyIntAlias = MyInt

func (mia *MyIntAlias) Get() int {
	return int(*mia)
}
