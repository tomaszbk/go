package runtime_test

import (
	"testing"
)

func TestTypelinksRace(t *testing.T) {
	output := runTestProg(t, "testprog", "TypelinksRace")
	want := "OK\n"
	if output != want {
		t.Fatalf("want %s, got %s\n", want, output)
	}
}
