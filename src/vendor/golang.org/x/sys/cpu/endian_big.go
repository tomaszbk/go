//go:build armbe || arm64be || m68k || mips || mips64 || mips64p32 || ppc || ppc64 || s390 || s390x || shbe || sparc || sparc64

package cpu

// IsBigEndian records whether the GOARCH's byte order is big endian.
const IsBigEndian = true
