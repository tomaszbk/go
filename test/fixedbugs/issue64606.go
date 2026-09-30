// build -race

//go:build race


package main

func main() {
	var o any = uint64(5)
	switch o.(type) {
	case int:
		goto ret
	case int8:
		goto ret
	case int16:
		goto ret
	case int32:
		goto ret
	case int64:
		goto ret
	case float32:
		goto ret
	case float64:
		goto ret
	default:
		goto ret
	}
ret:
}
