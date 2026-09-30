package main

import "fmt"

func init() {
	register("CheckAVX", CheckAVX)
}

func CheckAVX() {
	checkAVX()
	fmt.Println("OK")
}

func checkAVX()
