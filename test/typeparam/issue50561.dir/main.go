package main

import (
	"./diameter"
)

func main() {
	diameter.NewInboundHandler("hello", "world", "hi")
}
