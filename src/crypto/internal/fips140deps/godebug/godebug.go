package godebug

import (
	"internal/godebug"
)

type Setting godebug.Setting

func New(name string) *Setting {
	return (*Setting)(godebug.New(name))
}

func (s *Setting) Value() string {
	return (*godebug.Setting)(s).Value()
}

func Value(name string) string {
	return godebug.New(name).Value()
}
