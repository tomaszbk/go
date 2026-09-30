//go:build generate

package windows

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go eventlog.go service.go syscall_windows.go security_windows.go setupapi_windows.go
