package usedeprecated

import "testing"

func TestF(t *testing.T) {
	Legacy() // expect no deprecation notice.
	x()
}
