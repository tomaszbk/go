//go:build js && wasm

package os

// supportsCloseOnExec reports whether the platform supports the
// O_CLOEXEC flag.
const supportsCloseOnExec = false
