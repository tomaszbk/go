//go:build cmd_go_bootstrap || compiler_bootstrap

package telemetry

func MaybeParent()              {}
func MaybeChild()               {}
func Mode() string              { return "" }
func SetMode(mode string) error { return nil }
func Dir() string               { return "" }
