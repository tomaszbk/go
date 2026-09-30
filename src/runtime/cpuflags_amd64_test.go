package runtime_test

import (
	"runtime"
	"testing"
)

func TestHasAVX(t *testing.T) {
	t.Parallel()
	output := runTestProg(t, "testprog", "CheckAVX")
	ok := output == "OK\n"
	if *runtime.X86HasAVX != ok {
		t.Fatalf("x86HasAVX: %v, CheckAVX got:\n%s", *runtime.X86HasAVX, output)
	}
}
