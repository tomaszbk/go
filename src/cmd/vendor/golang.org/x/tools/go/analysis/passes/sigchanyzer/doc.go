// Package sigchanyzer defines an Analyzer that detects
// misuse of unbuffered signal as argument to signal.Notify.
//
// # Analyzer sigchanyzer
//
// sigchanyzer: check for unbuffered channel of os.Signal
//
// This checker reports call expression of the form
//
//	signal.Notify(c <-chan os.Signal, sig ...os.Signal),
//
// where c is an unbuffered channel, which can be at risk of missing the signal.
package sigchanyzer
