// Package profilerecord holds internal types used to represent profiling
// records with deep stack traces.
//
// TODO: Consider moving this to internal/runtime, see golang.org/issue/65355.
package profilerecord

type StackRecord struct {
	Stack []uintptr
}

type MemProfileRecord struct {
	ObjectSize                int64
	AllocObjects, FreeObjects int64
	Stack                     []uintptr
}

func (r *MemProfileRecord) InUseBytes() int64   { return r.InUseObjects() * r.ObjectSize }
func (r *MemProfileRecord) InUseObjects() int64 { return r.AllocObjects - r.FreeObjects }

type BlockProfileRecord struct {
	Count  int64
	Cycles int64
	Stack  []uintptr
}
