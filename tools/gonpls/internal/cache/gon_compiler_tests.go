package cache

import (
	"go/build/constraint"
	"path/filepath"
	"strings"
)

// isCompilerTest recognizes the standalone inputs to cmd/internal/testdir.
// These are not packages: neighboring files may deliberately redeclare main,
// shadow builtins, or contain errors. Limit this exception to the selected
// toolchain's test tree, and never split the multi-file fixtures in *.dir.
func isCompilerTest(goroot, filename string, src []byte) bool {
	if !isCompilerTestPath(goroot, filename) {
		return false
	}
	return hasCompilerTestRecipe(src)
}

func isCompilerTestPath(goroot, filename string) bool {
	if goroot == "" {
		return false
	}
	rel, err := filepath.Rel(filepath.Join(goroot, "test"), filename)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return false
	}
	for _, part := range strings.Split(rel, string(filepath.Separator)) {
		if strings.HasSuffix(part, ".dir") {
			return false
		}
	}
	return true
}

func hasCompilerTestRecipe(src []byte) bool {
	// Match the harness's first nonempty line, skipping build constraints.
	for _, line := range strings.Split(string(src), "\n") {
		if constraint.IsGoBuild(line) || constraint.IsPlusBuild(line) || strings.TrimSpace(line) == "" {
			continue
		}
		if !strings.HasPrefix(line, "//") {
			return false
		}
		fields := strings.Fields(strings.TrimPrefix(line, "//"))
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "run", "compile", "build", "buildrun", "runoutput", "asmcheck", "skip",
			"errorcheck", "errorcheckwithauto", "errorcheckoutput",
			"compiledir", "builddir", "buildrundir", "rundir", "runindir",
			"errorcheckdir", "errorcheckandrundir":
			return true
		}
		return false
	}
	return false
}
