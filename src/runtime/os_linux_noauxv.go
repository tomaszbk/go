//go:build linux && !arm && !arm64 && !loong64 && !mips && !mipsle && !mips64 && !mips64le && !s390x && !ppc64 && !ppc64le

package runtime

func archauxv(tag, val uintptr) {
}
