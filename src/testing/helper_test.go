package testing_test

import (
	"internal/testenv"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestTBHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		testTestHelper(t)

		// Check that calling Helper from inside a top-level test function
		// has no effect.
		t.Helper()
		t.Error("8")
		return
	}

	t.Parallel()

	cmd := testenv.Command(t, testenv.Executable(t), "-test.run=^TestTBHelper$")
	cmd = testenv.CleanCmdEnv(cmd)
	cmd.Env = append(cmd.Env, "GO_WANT_HELPER_PROCESS=1")
	out, _ := cmd.CombinedOutput()

	want := `--- FAIL: TestTBHelper \([^)]+\)
    helperfuncs_test.go:11: 0
    helperfuncs_test.go:43: 1
    helperfuncs_test.go:20: 2
    helperfuncs_test.go:45: 3
    helperfuncs_test.go:52: 4
    --- FAIL: TestTBHelper/sub \([^)]+\)
        helperfuncs_test.go:55: 5
        helperfuncs_test.go:20: 6
        helperfuncs_test.go:54: 7
    --- FAIL: TestTBHelper/sub2 \([^)]+\)
        helperfuncs_test.go:76: 11
    helperfuncs_test.go:80: recover 12
    helperfuncs_test.go:82: GenericFloat64
    helperfuncs_test.go:83: GenericInt
    helper_test.go:18: 8
    helperfuncs_test.go:69: 9
    helperfuncs_test.go:65: 10
`
	if !regexp.MustCompile(want).Match(out) {
		t.Errorf("got output:\n\n%s\nwant matching:\n\n%s", out, want)
	}
}

func TestTBHelperParallel(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		parallelTestHelper(t)
		return
	}

	t.Parallel()

	cmd := testenv.Command(t, testenv.Executable(t), "-test.run=^TestTBHelperParallel$")
	cmd = testenv.CleanCmdEnv(cmd)
	cmd.Env = append(cmd.Env, "GO_WANT_HELPER_PROCESS=1")
	out, _ := cmd.CombinedOutput()

	t.Logf("output:\n%s", out)

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	// We expect to see one "--- FAIL" line at the start
	// of the log, five lines of "parallel" logging,
	// and a final "FAIL" line at the end of the test.
	const wantLines = 7

	if len(lines) != wantLines {
		t.Fatalf("parallelTestHelper gave %d lines of output; want %d", len(lines), wantLines)
	}
	want := "helperfuncs_test.go:20: parallel"
	if got := strings.TrimSpace(lines[1]); got != want {
		t.Errorf("got second output line %q; want %q", got, want)
	}
}

// Issue 72794.
func TestHelperRange(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		rangeHelperHelper(t)
		return
	}

	t.Parallel()

	cmd := testenv.Command(t, testenv.Executable(t), "-test.run=^TestHelperRange$")
	cmd = testenv.CleanCmdEnv(cmd)
	cmd.Env = append(cmd.Env, "GO_WANT_HELPER_PROCESS=1")
	out, _ := cmd.CombinedOutput()
	want := `--- FAIL: TestHelperRange \([^)]+\)
    helperfuncs_test.go:135: range
    helperfuncs_test.go:135: range
`
	if !regexp.MustCompile(want).Match(out) {
		t.Errorf("got output:\n\n%s\nwant matching:\n\n%s", out, want)
	}
}

func BenchmarkTBHelper(b *testing.B) {
	f1 := func() {
		b.Helper()
	}
	f2 := func() {
		b.Helper()
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if i&1 == 0 {
			f1()
		} else {
			f2()
		}
	}
}
