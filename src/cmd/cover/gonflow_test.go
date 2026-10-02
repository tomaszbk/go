package main_test

import (
	"internal/testenv"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var gonFlowPaths = []string{"TestPresentOK", "TestAbsentOK", "TestPresentError", "TestAbsentError"}
var gonFlowTags = map[string][4]int{
	"global-handler":   {0, 0, 1, 1},
	"lambda-create":    {1, 1, 1, 1},
	"lambda-parse":     {1, 1, 1, 1},
	"lambda-increment": {1, 1, 0, 0},
	"lambda-call":      {1, 1, 1, 1},
	"lambda-after":     {1, 1, 1, 1},
	"chain-select":     {1, 1, 1, 1},
	"chain-after":      {1, 1, 0, 1},
	"coalesce-select":  {1, 1, 1, 1},
	"coalesce-after":   {1, 1, 1, 0},
	"assign-select":    {1, 1, 1, 1},
	"assign-after":     {1, 1, 1, 0},
}

func TestGonFlowCoverage(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	testenv.MustHaveExec(t)
	if coverageBaseline() == "" {
		t.Fatal("GON_BASELINE_GO must name an unmodified Go toolchain")
	}
	modes := errorFlowModes
	if testing.Short() {
		modes = []string{"count"}
	}
	for _, tc := range errorFlowToolchains(t) {
		for _, mode := range modes {
			t.Run(tc.name+"/"+mode, func(t *testing.T) {
				dir := tempDir(t)
				for from, to := range map[string]string{"common.go": "common.go", tc.variant + ".go": "flow.go", "flow_test.go": "flow_test.go"} {
					data, err := os.ReadFile(filepath.Join("testdata/gonflow", from))
					if err != nil {
						t.Fatal(err)
					}
					if err := os.WriteFile(filepath.Join(dir, to), data, 0666); err != nil {
						t.Fatal(err)
					}
				}
				if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module gonflow\n\ngo 1.26\n"), 0666); err != nil {
					t.Fatal(err)
				}
				exe := filepath.Join(dir, "flow.test.exe")
				run(tc.command(t, dir, "test", "-c", "-cover", "-covermode", mode, "-o", exe), t)
				source, err := os.ReadFile(filepath.Join(dir, "flow.go"))
				if err != nil {
					t.Fatal(err)
				}
				tags := map[string]errorFlowTag{}
				for i, line := range strings.Split(string(source), "\n") {
					if m := errorFlowTagRE.FindStringSubmatch(line); m != nil {
						tags[m[1]] = errorFlowTag{line: i + 1, col: len(line) - len(strings.TrimLeft(line, " \t")) + 1}
					}
				}
				for i, path := range gonFlowPaths {
					profile := filepath.Join(dir, path+".out")
					cmd := testenv.Command(t, exe, "-test.run=^"+path+"$", "-test.count=1", "-test.coverprofile="+profile)
					cmd.Dir = dir
					run(cmd, t)
					blocks := readProfileBlocks(t, profile, "flow.go", mode)
					for tag, wants := range gonFlowTags {
						loc, ok := tags[tag]
						if !ok {
							t.Fatalf("missing tag %s", tag)
						}
						count, found := 0, false
						for _, b := range blocks {
							if b.startLine <= loc.line && loc.line <= b.endLine && (b.startLine != loc.line || b.startCol <= loc.col) && (b.endLine != loc.line || loc.col < b.endCol) {
								count += b.count
								found = true
							}
						}
						if !found || count != wants[i] {
							t.Errorf("%s %s: count=%d found=%v want=%d", path, tag, count, found, wants[i])
						}
					}
				}
			})
		}
	}
}
