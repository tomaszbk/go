//go:build (goexperiment.boringcrypto && !boringcrypto) || (!goexperiment.boringcrypto && boringcrypto)

package boring_test

import "testing"

func TestNotBoring(t *testing.T) {
	t.Error("goexperiment.boringcrypto and boringcrypto should be equivalent build tags")
}
