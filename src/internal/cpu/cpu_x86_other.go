//go:build (386 || amd64) && (!darwin || ios) && !netbsd

package cpu

func osInit() {}
