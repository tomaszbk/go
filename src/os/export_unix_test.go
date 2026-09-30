//go:build unix || (js && wasm) || wasip1

package os

var SplitPath = splitPath
