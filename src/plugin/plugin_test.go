package plugin_test

import (
	_ "plugin"
	"testing"
)

func TestPlugin(t *testing.T) {
	// This test makes sure that executable that imports plugin
	// package can actually run. See issue #28789 for details.
}
