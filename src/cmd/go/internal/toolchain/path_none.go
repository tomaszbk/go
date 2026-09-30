//go:build !unix && !plan9 && !windows

package toolchain

import "io/fs"

// pathDirs returns the directories in the system search path.
func pathDirs() []string {
	return nil
}

// pathVersion returns the Go version implemented by the file
// described by de and info in directory dir.
// The analysis only uses the name itself; it does not run the program.
func pathVersion(dir string, de fs.DirEntry, info fs.FileInfo) (string, bool) {
	return "", false
}
