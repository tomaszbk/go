//go:build !unix && !windows

package ld

// Mmap allocates an in-heap output buffer with the given size. It copies
// any old data (if any) to the new buffer.
func (out *OutBuf) Mmap(filesize uint64) error {
	// We need space to put all the symbols before we apply relocations.
	oldheap := out.heap
	if filesize < uint64(len(oldheap)) {
		panic("mmap size too small")
	}
	out.heap = make([]byte, filesize)
	copy(out.heap, oldheap)
	return nil
}

func (out *OutBuf) munmap() { panic("unreachable") }
