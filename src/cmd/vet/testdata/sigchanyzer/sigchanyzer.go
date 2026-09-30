package sigchanyzer

import (
	"os"
	"os/signal"
)

func _() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt) // ERROR "misuse of unbuffered os.Signal channel as argument to signal.Notify"
}
