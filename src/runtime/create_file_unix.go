//go:build unix

package runtime

const canCreateFile = true

// create returns an fd to a write-only file.
func create(name *byte, perm int32) int32 {
	return open(name, _O_CREAT|_O_WRONLY|_O_TRUNC, perm)
}
