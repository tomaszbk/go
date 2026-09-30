package a

type Interaction[DataT InteractionDataConstraint] struct {
}

type InteractionDataConstraint interface {
	[]byte |
		UserCommandInteractionData
}

type UserCommandInteractionData struct {
	resolvedInteractionWithOptions
}

type resolvedInteractionWithOptions struct {
	Resolved Resolved `json:"resolved,omitempty"`
}

type Resolved struct {
	Users ResolvedData[User] `json:"users,omitempty"`
}

type ResolvedData[T ResolvedDataConstraint] map[uint64]T

type ResolvedDataConstraint interface {
	User | Message
}

type User struct{}

type Message struct {
	Interaction *Interaction[[]byte] `json:"interaction,omitempty"`
}
