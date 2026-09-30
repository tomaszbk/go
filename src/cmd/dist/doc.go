// Dist helps bootstrap, build, and test the Go distribution.
//
// Usage:
//
//	go tool dist [command]
//
// The commands are:
//
//	banner         print installation banner
//	bootstrap      rebuild everything
//	clean          deletes all built files
//	env [-p]       print environment (-p: include $PATH)
//	list [-json]   list all supported platforms
//	test [-h]      run Go test(s)
//	version        print Go version
package main
