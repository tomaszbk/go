package toolchain

import (
	"io/fs"
	"os"
	"path/filepath"

	"cmd/go/internal/gover"
)

// pathDirs returns the directories in the system search path.
func pathDirs() []string {
	return filepath.SplitList(os.Getenv("path"))
}

// pathVersion returns the Go version implemented by the file
// described by de and info in directory dir.
// The analysis only uses the name itself; it does not run the program.
func pathVersion(dir string, de fs.DirEntry, info fs.FileInfo) (string, bool) {
	v := gover.FromToolchain(de.Name())
	if v == "" || info.Mode()&0111 == 0 {
		return "", false
	}
	return v, true
}
