//go:build linux && (mips || mipsle || mips64 || mips64le)

package unix

type siErrnoCode struct {
	Code  int32
	Errno int32
}
