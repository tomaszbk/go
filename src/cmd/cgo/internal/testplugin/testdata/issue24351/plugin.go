package main

import "fmt"

func B(c chan bool) {
	go func() {
		fmt.Println(1.5)
		c <- true
	}()
}
