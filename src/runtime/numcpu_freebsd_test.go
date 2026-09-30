package runtime_test

import "testing"

func TestFreeBSDNumCPU(t *testing.T) {
	got := runTestProg(t, "testprog", "FreeBSDNumCPU")
	want := "OK\n"
	if got != want {
		t.Fatalf("expected %q, but got:\n%s", want, got)
	}
}
