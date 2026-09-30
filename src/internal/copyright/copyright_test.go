package copyright

import (
	"internal/testenv"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Gon removed the per-file copyright headers of Go Authors sources, so the
// root LICENSE is the notice that redistributions retain. It must stay intact.
func TestCopyright(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(testenv.GOROOT(t), "LICENSE"))
	if err != nil {
		t.Fatal(err)
	}
	license := strings.Join(strings.Fields(string(data)), " ")
	for _, want := range []string{
		"Copyright 2009 The Go Authors.",
		"Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.",
		"Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.",
		"Neither the name of Google LLC nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.",
		`THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"`,
	} {
		if !strings.Contains(license, want) {
			t.Errorf("LICENSE no longer contains %q", want)
		}
	}
}
