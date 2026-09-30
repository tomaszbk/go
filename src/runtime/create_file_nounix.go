//go:build !unix

package runtime

const canCreateFile = false

func create(name *byte, perm int32) int32 {
	throw("unimplemented")
	return -1
}
