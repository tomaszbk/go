package os_test

import (
	"os"
	"testing"
)

func init() {
	pipeDeadlinesTestCases = []pipeDeadlineTest{
		{
			"named overlapped pipe",
			func(t *testing.T) (r, w *os.File) {
				name := pipeName()
				w = newBytePipe(t, name, true)
				r = newFileOverlapped(t, name, true)
				return
			},
		},
	}
}
