package main

import "test/a"

func main() {
	stop := start()
	defer stop()
}

func start() func() {
	return a.Start().Stop
}
