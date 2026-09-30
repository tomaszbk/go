package b

import (
	"./a"
)

// InteractionRequest is an incoming request Interaction
type InteractionRequest[T a.InteractionDataConstraint] struct {
	a.Interaction[T]
}
