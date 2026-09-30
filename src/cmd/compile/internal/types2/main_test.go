package types2_test

import (
	"go/build"
	"internal/testenv"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	build.Default.GOROOT = testenv.GOROOT(nil)
	os.Exit(m.Run())
}
