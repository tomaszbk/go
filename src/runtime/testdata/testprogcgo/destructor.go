package main

// extern void registerDestructor();
import "C"

import "fmt"

func init() {
	register("DestructorCallback", DestructorCallback)
}

//export GoDestructorCallback
func GoDestructorCallback() {
}

func DestructorCallback() {
	C.registerDestructor()
	fmt.Println("OK")
}
