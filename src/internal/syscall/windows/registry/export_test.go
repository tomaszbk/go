//go:build windows

package registry

func (k Key) SetValue(name string, valtype uint32, data []byte) error {
	return k.setValue(name, valtype, data)
}
