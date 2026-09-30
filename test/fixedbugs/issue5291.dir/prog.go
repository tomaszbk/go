package main

import (
	"./pkg1"
)

type message struct { // Presence of this creates a crash
	data pkg1.Data
}

func main() {
	pkg1.CrashCall()
}
