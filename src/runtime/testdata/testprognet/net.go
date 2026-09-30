package main

import (
	"fmt"
	"net"
)

func init() {
	registerInit("NetpollDeadlock", NetpollDeadlockInit)
	register("NetpollDeadlock", NetpollDeadlock)
}

func NetpollDeadlockInit() {
	fmt.Println("dialing")
	c, err := net.Dial("tcp", "localhost:14356")
	if err == nil {
		c.Close()
	} else {
		fmt.Println("error: ", err)
	}
}

func NetpollDeadlock() {
	fmt.Println("done")
}
