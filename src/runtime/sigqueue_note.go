// The current implementation of notes on Darwin is not async-signal-safe,
// so on Darwin the sigqueue code uses different functions to wake up the
// signal_recv thread. This file holds the non-Darwin implementations of
// those functions. These functions will never be called.

//go:build !darwin && !plan9

package runtime

func sigNoteSetup(*note) {
	throw("sigNoteSetup")
}

func sigNoteSleep(*note) {
	throw("sigNoteSleep")
}

func sigNoteWakeup(*note) {
	throw("sigNoteWakeup")
}
