package p

import (
	"./b"
)

// ResponseWriterMock mocks corde's ResponseWriter interface
type ResponseWriterMock struct {
	x b.InteractionRequest[[]byte]
}
