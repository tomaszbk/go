package ssabase

// A Register is a machine register, like AX.
// They are numbered densely from 0 (for each architecture).
type Register struct {
	Num    int32 // dense numbering
	ObjNum int16 // register number from cmd/internal/obj/$ARCH
	Name   string
}

func (r *Register) String() string {
	return r.Name
}
