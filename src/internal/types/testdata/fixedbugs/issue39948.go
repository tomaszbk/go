package p

type T[P any] interface{
	P // ERROR "term cannot be a type parameter"
}
