//go:build boringcrypto

package main_test

import (
	"os"
	"testing"
)

func TestBoringInternalLink(t *testing.T) {
	tg := testgo(t)
	defer tg.cleanup()
	tg.parallel()
	tg.tempFile("main.go", `package main
		import "crypto/sha1"
		func main() {
			sha1.New()
		}`)
	tg.run("build", "-ldflags=-w -extld=false", "-o", os.DevNull, tg.path("main.go"))
	tg.run("build", "-ldflags=-extld=false", "-o", os.DevNull, tg.path("main.go"))
}
