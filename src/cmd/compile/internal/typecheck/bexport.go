package typecheck

// Tags. Must be < 0.
const (
	// Objects
	packageTag = -(iota + 1)
	constTag
	typeTag
	varTag
	funcTag
	endTag
)
