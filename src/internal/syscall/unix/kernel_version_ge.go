package unix

// KernelVersionGE checks if the running kernel version
// is greater than or equal to the provided version.
func KernelVersionGE(x, y int) bool {
	xx, yy := KernelVersion()

	return xx > x || (xx == x && yy >= y)
}
