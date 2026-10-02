// Package testconditional checks lazy conditional expressions in cgo source.
// The legacy program runs with both Gon and the unmodified GON_BASELINE_GO;
// the modern program runs with Gon and must produce identical output.
package testconditional

import (
	"bytes"
	"encoding/json"
	"internal/testenv"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// newModule returns a directory holding a module made of the shared test files
// and, if program is not empty, the named testdata file as prog.go.
func newModule(t *testing.T, program string) string {
	t.Helper()
	dir := t.TempDir()
	write := func(name string, data []byte) {
		if err := os.WriteFile(filepath.Join(dir, name), data, 0o666); err != nil {
			t.Fatal(err)
		}
	}
	copyFile := func(src, dst string) {
		data, err := os.ReadFile(filepath.Join("testdata", src))
		if err != nil {
			t.Fatal(err)
		}
		write(dst, data)
	}
	write("go.mod", []byte("module conditionalcgo\n\ngo 1.26\n"))
	for _, name := range []string{"common.go", "clib.h", "clib.c"} {
		copyFile(name, name)
	}
	if program != "" {
		copyFile(program, "prog.go")
	}
	return dir
}

// goCommand runs goTool in dir. A baseline toolchain is run without the
// environment variables that would point it at this repository's GOROOT.
func goCommand(t *testing.T, goTool string, baseline bool, dir string, args ...string) (stdout, stderr string, err error) {
	t.Helper()
	cmd := exec.Command(goTool, args...)
	cmd.Dir = dir
	for _, entry := range os.Environ() {
		if baseline && (strings.HasPrefix(entry, "GOROOT=") || strings.HasPrefix(entry, "GOTOOLDIR=")) {
			continue
		}
		cmd.Env = append(cmd.Env, entry)
	}
	cmd.Env = append(cmd.Env, "GOENV=off", "GOFLAGS=", "GOTOOLCHAIN=local", "GO111MODULE=on")
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err = cmd.Run()
	return out.String(), errOut.String(), err
}

func runProgram(t *testing.T, goTool string, baseline bool, program string) string {
	t.Helper()
	dir := newModule(t, program)
	stdout, stderr, err := goCommand(t, goTool, baseline, dir, "run", ".")
	if err != nil {
		t.Fatalf("%s (baseline=%v): go run: %v\nstdout:\n%s\nstderr:\n%s", program, baseline, err, stdout, stderr)
	}
	if !strings.HasSuffix(stdout, "\nPASS\n") {
		t.Fatalf("%s (baseline=%v): program did not report PASS:\n%s", program, baseline, stdout)
	}
	return stdout
}

func TestPairedCgoConditional(t *testing.T) {
	testenv.MustHaveGoRun(t)
	testenv.MustHaveCGO(t)
	goTool := testenv.GoToolPath(t)

	baseline := os.Getenv("GON_BASELINE_GO")
	if baseline == "" {
		t.Fatal("GON_BASELINE_GO must name an unmodified Go toolchain for the executable pair")
	}

	legacy := runProgram(t, goTool, false, "legacy.go")
	modern := runProgram(t, goTool, false, "modern.go")
	if legacy != modern {
		t.Errorf("legacy and modern programs behave differently\nlegacy:\n%s\nmodern:\n%s", legacy, modern)
	}

	t.Run("vet", func(t *testing.T) {
		dir := newModule(t, "modern.go")
		if stdout, stderr, err := goCommand(t, goTool, false, dir, "vet", "."); err != nil {
			t.Errorf("go vet: %v\nstdout:\n%s\nstderr:\n%s", err, stdout, stderr)
		}
	})

	upstream := runProgram(t, baseline, true, "legacy.go")
	if legacy != upstream {
		t.Errorf("legacy program behaves differently with the baseline toolchain %s\nfork:\n%s\nbaseline:\n%s", baseline, legacy, upstream)
	} else {
		t.Logf("legacy program output is identical with the baseline toolchain %s", baseline)
	}
}

// Compile the maintained cgo sources against upstream go/ast, then process
// ordinary Go source with that binary. This exercises the bootstrap adapters
// without rebuilding unrelated compiler packages or replacing installed tools.
func TestCgoConditionalBootstrap(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	testenv.MustHaveCGO(t)
	baseline := os.Getenv("GON_BASELINE_GO")
	if baseline == "" {
		t.Fatal("GON_BASELINE_GO must name an unmodified Go toolchain")
	}
	dir := t.TempDir()
	stdout, stderr, err := goCommand(t, baseline, true, dir, "env", "GOROOT")
	if err != nil {
		t.Fatalf("baseline GOROOT: %v\n%s", err, stderr)
	}
	baselineDir := filepath.Join(strings.TrimSpace(stdout), "src", "cmd", "cgo")
	sourceDir := filepath.Join(testenv.GOROOT(t), "src", "cmd", "cgo")
	replace := make(map[string]string)
	for _, pair := range []struct {
		dir        string
		maintained bool
	}{{baselineDir, false}, {sourceDir, true}} {
		files, err := filepath.Glob(filepath.Join(pair.dir, "*.go"))
		if err != nil || len(files) == 0 {
			t.Fatalf("cgo sources in %s: %v (files=%d)", pair.dir, err, len(files))
		}
		for _, file := range files {
			replacement := ""
			if pair.maintained {
				replacement = file
			}
			replace[filepath.Join(baselineDir, filepath.Base(file))] = replacement
		}
	}
	data, err := json.Marshal(struct{ Replace map[string]string }{replace})
	if err != nil {
		t.Fatal(err)
	}
	overlay := filepath.Join(dir, "overlay.json")
	if err := os.WriteFile(overlay, data, 0o666); err != nil {
		t.Fatal(err)
	}
	cgo := filepath.Join(dir, "cgo.exe")
	stdout, stderr, err = goCommand(t, baseline, true, dir, "build", "-overlay", overlay, "-tags=compiler_bootstrap", "-o", cgo, "cmd/cgo")
	if err != nil {
		t.Fatalf("bootstrap cgo build: %v\n%s%s", err, stdout, stderr)
	}
	source := filepath.Join(dir, "legacy.go")
	if err := os.WriteFile(source, []byte("package p\n// static int answer(void) { return 42; }\nimport \"C\"\nfunc answer() int { return int(C.answer()) }\n"), 0o666); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(cgo, "-objdir", dir, source)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("bootstrap cgo processing: %v\n%s", err, out)
	}
	if _, err := os.Stat(filepath.Join(dir, "legacy.cgo1.go")); err != nil {
		t.Fatalf("bootstrap cgo output: %v", err)
	}
}

const invalidPrologue = `package main
/*
#include <errno.h>
typedef struct { int *p; int n; } c_buf;
static int c_sum(c_buf *b, int n) { return (b ? b->n : 0) + n; }
static int c_div(int a, int b) { if (!b) { errno = EDOM; return -1; } return a / b; }
*/
import "C"
import "strconv"
var _ = strconv.Atoi
func main() { _ = f }
`

func TestCgoConditionalDiagnostics(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	testenv.MustHaveCGO(t)
	goTool := testenv.GoToolPath(t)
	tests := []struct{ name, source, want string }{
		{"control_nil_target", `func f(c bool) int { return int(C.c_sum((if c { nil } else { nil }), 1)) }`, ""},
		{"control_branch_c_refs", `func f(c bool) C.int { return if C.c_div(1, 1) > 0 { C.int(1) } else { C.int(2) } }`, ""},
		{"control_nested_function_boundary", `func f(c bool) int { var b C.c_buf; return int(C.c_sum(&b, C.int(func() int { return if c { strconv.Atoi("1") or err { return -1 } } else { 2 } }()))) }`, ""},
		{"missing_else", `func f(c bool) int { return int(if c { C.c_div(1, 1) }) }`, `conditional expression requires an else branch`},
		{"multivalue_branch", `func f(c bool) int { n := if c { strconv.Atoi("1") } else { 2 }; return n }`, `multiple-value`},
		{"mismatched_c_types", `func f(c bool) int { n := if c { C.int(1) } else { C.long(2) }; return int(n) }`, `mismatched types`},
		{"pointer_condition_propagation", `func f(c bool) (int, error) { var b C.c_buf; n := C.c_sum(if strconv.Atoi("1")! > 0 { &b } else { nil }, 1); return int(n), nil }`, `cannot use error propagation or an "or" handler in an argument of a C call that cgo must rewrite`},
		{"pointer_branch_propagation", `func f(c bool) (int, error) { var b C.c_buf; n := C.c_sum(&b, C.int(if c { strconv.Atoi("1")! } else { 2 })); return int(n), nil }`, `cannot use error propagation or an "or" handler in an argument of a C call that cgo must rewrite`},
		{"pointer_branch_handler", `func f(c bool) (int, error) { var b C.c_buf; n := C.c_sum(&b, C.int(if c { strconv.Atoi("1") or err { return 0, err } } else { 2 })); return int(n), nil }`, `cannot use error propagation or an "or" handler in an argument of a C call that cgo must rewrite`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dir := t.TempDir()
			for name, contents := range map[string]string{"go.mod": "module invalid\n\ngo 1.26\n", "invalid.go": invalidPrologue + test.source + "\n"} {
				if err := os.WriteFile(filepath.Join(dir, name), []byte(contents), 0o666); err != nil {
					t.Fatal(err)
				}
			}
			stdout, stderr, err := goCommand(t, goTool, false, dir, "build", "-o", filepath.Join(dir, "invalid.exe"), ".")
			out := stdout + stderr
			if test.want == "" {
				if err != nil {
					t.Fatalf("valid program failed to build: %v\n%s", err, out)
				}
				return
			}
			if err == nil {
				t.Fatal("invalid program built successfully")
			}
			for _, crash := range []string{"panic:", "goroutine ", "unexpected type", "internal compiler error"} {
				if strings.Contains(out, crash) {
					t.Fatalf("invalid program crashed a tool (%q):\n%s", crash, out)
				}
			}
			re := regexp.MustCompile(`invalid\.go:\d+:\d+: .*` + test.want)
			if !re.MatchString(out) {
				t.Fatalf("output does not report %q at invalid.go:\n%s", test.want, out)
			}
		})
	}
}
