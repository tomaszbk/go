//go:build plan9

package socktest

// Sockets maps a socket descriptor to the status of socket.
type Sockets map[int]Status

func familyString(family int) string { return "<nil>" }

func typeString(sotype int) string { return "<nil>" }

func protocolString(proto int) string { return "<nil>" }
