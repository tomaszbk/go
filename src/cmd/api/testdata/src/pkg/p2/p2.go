package p2

type Twoer interface {
	// Deprecated: No good.
	PackageTwoMeth()
}

// Deprecated: No good.
func F() string {}

func G() Twoer {}

func NewError(s string) error {}
