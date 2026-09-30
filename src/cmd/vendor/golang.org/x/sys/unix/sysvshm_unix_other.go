//go:build (darwin && !ios) || zos

package unix

// SysvShmCtl performs control operations on the shared memory segment
// specified by id.
func SysvShmCtl(id, cmd int, desc *SysvShmDesc) (result int, err error) {
	return shmctl(id, cmd, desc)
}
