// compile


package p

type hasPtrs struct {
        x [2]*int
	// Note: array size needs to be >1 to force this type to be not SSAable.
	// The bug triggers only for OpMove, which is only used for unSSAable types.
}

func main() {
        var x *hasPtrs       // Can be local, global, or arg; nil or non-nil.
        var y *hasPtrs = nil // Must initialize to nil.
        *x = *y
}
