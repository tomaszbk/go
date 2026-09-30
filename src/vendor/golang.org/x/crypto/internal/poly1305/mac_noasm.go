//go:build (!amd64 && !loong64 && !ppc64le && !ppc64 && !riscv64 && !s390x) || !gc || purego

package poly1305

type mac struct{ macGeneric }
