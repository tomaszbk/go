package multipart

import (
	"net/textproto"
	_ "unsafe" // for go:linkname
)

// readMIMEHeader is defined in package [net/textproto].
//
//go:linkname readMIMEHeader net/textproto.readMIMEHeader
func readMIMEHeader(r *textproto.Reader, maxMemory, maxHeaders int64) (textproto.MIMEHeader, error)
