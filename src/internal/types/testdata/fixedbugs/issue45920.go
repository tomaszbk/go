package p

func f1[T any, C chan T | <-chan T](ch C) {}

func _(ch chan int)   { f1(ch) }
func _(ch <-chan int) { f1(ch) }
func _(ch chan<- int) { f1 /* ERROR "chan<- int does not satisfy chan int | <-chan int" */ (ch) }

func f2[T any, C chan T | chan<- T](ch C) {}

func _(ch chan int)   { f2(ch) }
func _(ch <-chan int) { f2 /* ERROR "<-chan int does not satisfy chan int | chan<- int" */ (ch) }
func _(ch chan<- int) { f2(ch) }
