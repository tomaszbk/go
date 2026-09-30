//go:build (js && wasm) || wasip1

package http

import (
	"context"
	"net"
)

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return nil
}
