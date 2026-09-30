package p

type S1 struct{}
type S2 struct{}

func _[P *S1|*S2]() {
	_= []P{{ /* ERROR "invalid composite literal element type P (no common underlying type)" */ }}
}

func _[P *S1|S1]() {
	_= []P{{ /* ERROR "invalid composite literal element type P (no common underlying type)" */ }}
}
