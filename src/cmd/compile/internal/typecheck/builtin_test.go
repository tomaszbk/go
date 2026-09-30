package typecheck

import (
	"bytes"
	"internal/testenv"
	"os"
	"testing"
)

func TestBuiltin(t *testing.T) {
	testenv.MustHaveGoRun(t)
	t.Parallel()

	old, err := os.ReadFile("builtin.go")
	if err != nil {
		t.Fatal(err)
	}

	new, err := testenv.Command(t, testenv.GoToolPath(t), "run", "mkbuiltin.go", "-stdout").Output()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(old, new) {
		t.Fatal("builtin.go out of date; run mkbuiltin.go")
	}
}
