//go:build generate

package windows

//go:generate go run ../../../syscall/mksyscall_windows.go -output zsyscall_windows.go syscall_windows.go security_windows.go psapi_windows.go symlink_windows.go version_windows.go
