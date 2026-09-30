//go:build js || wasip1 || windows

package net

const readMsgFlags = 0

func setReadMsgCloseOnExec(oob []byte) {}
