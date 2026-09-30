//go:build generate

package registry

//go:generate go run ../../../../syscall/mksyscall_windows.go -output zsyscall_windows.go syscall.go
