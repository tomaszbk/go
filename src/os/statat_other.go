//go:build (js && wasm) || plan9

package os

func (f *File) lstatatNolog(name string) (FileInfo, error) {
	// These platforms don't have fstatat, so use stat instead.
	return Lstat(f.name + "/" + name)
}
