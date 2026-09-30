package time_test

import (
	"testing"
	. "time"
)

func TestAndroidTzdata(t *testing.T) {
	undo := ForceAndroidTzdataForTest()
	defer undo()
	if _, err := LoadLocation("America/Los_Angeles"); err != nil {
		t.Error(err)
	}
}
