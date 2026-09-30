//go:build wasip1

package syscall

func JoinPath(dir, file string) string {
	return joinPath(dir, file)
}
