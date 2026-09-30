package main

func main[T /* ERROR "func main must have no type parameters" */ any]() {}
