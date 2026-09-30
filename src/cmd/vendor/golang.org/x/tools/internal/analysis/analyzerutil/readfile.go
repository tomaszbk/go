package analyzerutil

// This file defines helpers for calling [analysis.Pass.ReadFile].

import (
	"go/token"
	"os"

	"golang.org/x/tools/go/analysis"
)

// ReadFile reads a file and adds it to the FileSet in pass
// so that we can report errors against it using lineStart.
func ReadFile(pass *analysis.Pass, filename string) ([]byte, *token.File, error) {
	readFile := pass.ReadFile
	if readFile == nil {
		readFile = os.ReadFile
	}
	content, err := readFile(filename)
	if err != nil {
		return nil, nil, err
	}
	tf := pass.Fset.AddFile(filename, -1, len(content))
	tf.SetLinesForContent(content)
	return content, tf, nil
}
