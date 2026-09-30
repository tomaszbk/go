// errorcheck


package main

func f(x int, y ...int) // ok

func h(x, y ...int) // ERROR "[.][.][.]"

func i(x int, y ...int, z float32) // ERROR "[.][.][.]"
