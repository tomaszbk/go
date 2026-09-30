package a

func F[T interface{ chan int }](c T) {}
