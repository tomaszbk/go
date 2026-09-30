// Export guts for testing.
// Since testing imports os and os imports internal/poll,
// the internal/poll tests can not be in package poll.

package poll

var Consume = consume

type XFDMutex struct {
	fdMutex
}

func (mu *XFDMutex) Incref() bool {
	return mu.incref()
}

func (mu *XFDMutex) IncrefAndClose() bool {
	return mu.increfAndClose()
}

func (mu *XFDMutex) Decref() bool {
	return mu.decref()
}

func (mu *XFDMutex) RWLock(read bool) bool {
	return mu.rwlock(read, waitLock)
}

func (mu *XFDMutex) RWUnlock(read bool) bool {
	return mu.rwunlock(read)
}
