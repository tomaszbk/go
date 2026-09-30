//go:build linux && !(mips || mipsle || mips64 || mips64le)

package unix

type siErrnoCode struct {
	Errno int32
	Code  int32
}
