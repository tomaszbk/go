//go:build 386 || amd64p32 || (amd64 && ((!darwin && !netbsd) || !gc))

package cpu

func darwinSupportsAVX512() bool {
	panic("only implemented for gc && amd64 && darwin")
}
