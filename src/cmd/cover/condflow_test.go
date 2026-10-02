package main_test

import (
	"bytes"
	"internal/testenv"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var condFlowPaths = []string{"TestThenOK", "TestElseOK", "TestThenError", "TestElseError"}

// Conditional expressions count as part of their enclosing statement, rather
// than assigning a counter to each value branch. Handler bodies and function
// literals still have independent counters. A propagating error in either
// branch must leave the statements after the expression uncovered.
var condFlowTags = map[string][4]int{
	"select":             {2, 2, 2, 2},
	"select-after":       {2, 2, 2, 2},
	"prop-defer":         {1, 1, 1, 1},
	"prop-select":        {1, 1, 1, 1},
	"prop-after":         {1, 1, 0, 0},
	"handle-select":      {1, 1, 1, 1},
	"handle-then":        {0, 0, 1, 0},
	"handle-else":        {0, 0, 0, 1},
	"handle-after":       {1, 1, 0, 0},
	"closure-then-parse": {1, 0, 1, 0},
	"closure-then-after": {1, 0, 0, 0},
	"closure-else-parse": {0, 1, 0, 1},
	"closure-else-after": {0, 1, 0, 0},
	"closure-call":       {1, 1, 1, 1},
	"closure-fail":       {0, 0, 1, 1},
	"closure-after":      {1, 1, 0, 0},
	"nested":             {1, 1, 1, 1},
	"nested-after":       {1, 1, 1, 1},
}

// TestCondFlowCoverage runs the same assertions against legacy Go and Gon
// spelling and checks their profiles in every mode. GON_BASELINE_GO also runs
// the legacy version with an unmodified toolchain.
func TestCondFlowCoverage(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	testenv.MustHaveExec(t)
	t.Parallel()
	if coverageBaseline() == "" {
		t.Log("GON_BASELINE_GO is not set; not running the legacy version with an unmodified toolchain")
	}
	modes := errorFlowModes
	if testing.Short() {
		modes = []string{"count"}
	}
	for _, tc := range errorFlowToolchains(t) {
		for _, mode := range modes {
			t.Run(tc.name+"/"+mode, func(t *testing.T) {
				t.Parallel()
				testCondFlowTests(t, tc, mode)
			})
			t.Run(tc.name+"/"+mode+"/build", func(t *testing.T) {
				t.Parallel()
				testCondFlowBuild(t, tc, mode)
			})
		}
	}
}

func copyCondFlow(t *testing.T, dir, module, subdir, variant string, withTests bool) {
	t.Helper()
	pkgDir := filepath.Join(dir, subdir)
	if err := os.MkdirAll(pkgDir, 0777); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{"common.go": "common.go", variant + ".go": "flow.go"}
	if withTests {
		files["flow_test.go"] = "flow_test.go"
	}
	for from, to := range files {
		data, err := os.ReadFile(filepath.Join("testdata/condflow", from))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pkgDir, to), data, 0666); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module "+module+"\n\ngo 1.26\n"), 0666); err != nil {
		t.Fatal(err)
	}
}

func testCondFlowTests(t *testing.T, tc errorFlowToolchain, mode string) {
	dir := tempDir(t)
	copyCondFlow(t, dir, "condflow", ".", tc.variant, true)
	exe := filepath.Join(dir, "condflow.test.exe")
	run(tc.command(t, dir, "test", "-c", "-cover", "-covermode", mode, "-o", exe), t)
	tags := readCondFlowTags(t, filepath.Join(dir, "flow.go"))
	for i, path := range condFlowPaths {
		profile := filepath.Join(dir, path+".out")
		cmd := testenv.Command(t, exe, "-test.run=^"+path+"$", "-test.v", "-test.count=1", "-test.coverprofile="+profile)
		cmd.Dir = dir
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%s: %v\n%s", path, err, out)
		}
		if !bytes.Contains(out, []byte("--- PASS: "+path+" ")) {
			t.Fatalf("%s did not run or did not pass:\n%s", path, out)
		}
		checkCondFlowProfile(t, profile, mode, i, tags)
	}
}

// The build path exercises the other coverage-data collection mechanism,
// including an error inside a chosen expression branch.
func testCondFlowBuild(t *testing.T, tc errorFlowToolchain, mode string) {
	dir := tempDir(t)
	copyCondFlow(t, dir, "condflowmain", "flow", tc.variant, false)
	const main = `package main
import "condflowmain/flow"
func main() {
	flow.Select(true)
	flow.Select(true)
	flow.Propagate(true, "x", "7")
	flow.Handle(true, "x", "7")
	flow.Closure(true, "x", "7")
	flow.Nested(true)
}
`
	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(main), 0666); err != nil {
		t.Fatal(err)
	}
	exe := filepath.Join(dir, "condflowmain.exe")
	run(tc.command(t, dir, "build", "-cover", "-covermode", mode, "-o", exe), t)
	covdir := filepath.Join(dir, "covdata")
	if err := os.Mkdir(covdir, 0777); err != nil {
		t.Fatal(err)
	}
	cmd := testenv.Command(t, exe)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+covdir)
	run(cmd, t)
	profile := filepath.Join(dir, "profile.txt")
	run(tc.command(t, dir, "tool", "covdata", "textfmt", "-i="+covdir, "-o="+profile), t)
	tags := readCondFlowTags(t, filepath.Join(dir, "flow", "flow.go"))
	checkCondFlowProfile(t, profile, mode, 2, tags)
}

func readCondFlowTags(t *testing.T, file string) map[string]errorFlowTag {
	t.Helper()
	src, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	tags := make(map[string]errorFlowTag)
	for i, line := range strings.Split(string(src), "\n") {
		m := errorFlowTagRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		if _, dup := tags[m[1]]; dup {
			t.Fatalf("%s: duplicate tag %q", file, m[1])
		}
		if _, ok := condFlowTags[m[1]]; !ok {
			t.Errorf("%s: unexpected tag %q", file, m[1])
		}
		tags[m[1]] = errorFlowTag{line: i + 1, col: len(line) - len(strings.TrimLeft(line, " \t")) + 1}
	}
	for tag := range condFlowTags {
		if _, ok := tags[tag]; !ok {
			t.Errorf("%s: missing tag %q", file, tag)
		}
	}
	return tags
}

func checkCondFlowProfile(t *testing.T, profile, mode string, path int, tags map[string]errorFlowTag) {
	t.Helper()
	blocks := readProfileBlocks(t, profile, "flow.go", mode)
	if len(blocks) == 0 {
		t.Fatalf("%s: no blocks for flow.go", profile)
	}
	for i := 1; i < len(blocks); i++ {
		a, b := blocks[i-1], blocks[i]
		if a.endLine > b.startLine || a.endLine == b.startLine && a.endCol > b.startCol {
			t.Errorf("overlapping blocks %v and %v", a, b)
		}
	}
	for tag, pos := range tags {
		got := -1
		for _, b := range blocks {
			if (b.startLine < pos.line || b.startLine == pos.line && b.startCol <= pos.col) &&
				(pos.line < b.endLine || pos.line == b.endLine && pos.col < b.endCol) {
				got = b.count
				break
			}
		}
		want := condFlowTags[tag][path]
		if got < 0 {
			t.Errorf("%s: statement %q at line %d is in no block", condFlowPaths[path], tag, pos.line)
		} else if mode == "set" {
			if (got > 0) != (want > 0) {
				t.Errorf("%s: statement %q has count %d, want executed=%v", condFlowPaths[path], tag, got, want > 0)
			}
		} else if got != want {
			t.Errorf("%s: statement %q has count %d, want %d", condFlowPaths[path], tag, got, want)
		}
	}
}
