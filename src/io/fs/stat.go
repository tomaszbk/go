package fs

// A StatFS is a file system with a Stat method.
type StatFS interface {
	FS

	// Stat returns a FileInfo describing the file.
	// If there is an error, it should be of type *PathError.
	Stat(name string) (FileInfo, error)
}

// Stat returns a [FileInfo] describing the named file from the file system.
//
// If fsys implements [StatFS], Stat calls fsys.Stat.
// Otherwise, Stat opens the [File] to stat it.
func Stat(fsys FS, name string) (FileInfo, error) {
	if fsys, ok := fsys.(StatFS); ok {
		return fsys.Stat(name)
	}

	file, err := fsys.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file.Stat()
}
