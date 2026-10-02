package testconditional

import (
	"internal/testenv"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func gonFeatureModule(t *testing.T, feature, variant string) string {
	t.Helper()
	dir := t.TempDir()
	src, err := os.ReadFile(filepath.Join("testdata", feature+"_"+variant+".go"))
	if err != nil {
		t.Fatal(err)
	}
	for name, data := range map[string][]byte{
		"go.mod":       []byte("module gonfeaturescgo\n\ngo 1.26\n"),
		"main.go":      src,
		"main_test.go": []byte("package main\nimport \"testing\"\nfunc TestScenario(t *testing.T) { main() }\n"),
	} {
		if err := os.WriteFile(filepath.Join(dir, name), data, 0666); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestPairedCgoGonFeatures(t *testing.T) {
	testenv.MustHaveGoRun(t)
	testenv.MustHaveCGO(t)
	baseline := os.Getenv("GON_BASELINE_GO")
	if baseline == "" {
		t.Fatal("GON_BASELINE_GO must name an unmodified Go toolchain")
	}
	for _, feature := range []string{"lambda", "nullsafety"} {
		t.Run(feature, func(t *testing.T) {
			var output string
			for _, tc := range []struct {
				variant, tool string
				baseline      bool
			}{
				{"legacy", testenv.GoToolPath(t), false},
				{"modern", testenv.GoToolPath(t), false},
				{"legacy", baseline, true},
			} {
				dir := gonFeatureModule(t, feature, tc.variant)
				out, stderr, err := goCommand(t, tc.tool, tc.baseline, dir, "run", ".")
				if err != nil {
					t.Fatalf("%s baseline=%v: %v\n%s%s", tc.variant, tc.baseline, err, out, stderr)
				}
				if !strings.HasSuffix(out, "PASS\n") {
					t.Fatalf("scenario did not pass: %s", out)
				}
				if output == "" {
					output = out
				} else if output != out {
					t.Fatalf("unequal behavior: want %q, got %q", output, out)
				}
				out, stderr, err = goCommand(t, tc.tool, tc.baseline, dir, "test", "-cover", "-count=1", ".")
				if err != nil {
					t.Fatalf("%s baseline=%v coverage: %v\n%s%s", tc.variant, tc.baseline, err, out, stderr)
				}
				if !tc.baseline {
					out, stderr, err = goCommand(t, tc.tool, false, dir, "vet", ".")
					if err != nil {
						t.Fatalf("%s vet: %v\n%s%s", tc.variant, err, out, stderr)
					}
				}
			}
		})
	}
}
