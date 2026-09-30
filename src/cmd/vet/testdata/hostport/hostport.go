// This file contains tests for the hostport checker.

package hostport

import (
	"fmt"
	"net"
)

func _(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port) // ERROR "address format .%s:%d. does not work with IPv6"
	net.Dial("tcp", addr)
}
