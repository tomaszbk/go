//go:build (freebsd && amd64) || (freebsd && arm64) || (freebsd && riscv64)

package syscall

import "unsafe"

func (any *anyMessage) parseRouteMessage(b []byte) *RouteMessage {
	p := (*RouteMessage)(unsafe.Pointer(any))
	return &RouteMessage{Header: p.Header, Data: b[rsaAlignOf(int(unsafe.Offsetof(p.Header.Rmx))+SizeofRtMetrics):any.Msglen]}
}

func (any *anyMessage) parseInterfaceMessage(b []byte) *InterfaceMessage {
	p := (*InterfaceMessage)(unsafe.Pointer(any))
	return &InterfaceMessage{Header: p.Header, Data: b[int(unsafe.Offsetof(p.Header.Data))+int(p.Header.Data.Datalen) : any.Msglen]}
}
