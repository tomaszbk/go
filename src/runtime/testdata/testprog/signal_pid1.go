package main

import (
	"fmt"
	"os"
	"time"
)

func init() {
	register("SignalPid1", SignalPid1)
}

// SignalPid1 is a helper for TestSignalPid1.
func SignalPid1() {
	if os.Getpid() != 1 {
		fmt.Fprintln(os.Stderr, "I am not PID 1")
		return
	}
	fmt.Println("ready")

	time.Sleep(time.Hour)
}
