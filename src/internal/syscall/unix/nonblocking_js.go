//go:build js && wasm

package unix

func IsNonblock(fd int) (nonblocking bool, err error) {
	return false, nil
}

func HasNonblockFlag(flag int) bool {
	return false
}
