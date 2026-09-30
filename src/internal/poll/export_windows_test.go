// Export guts for testing on windows.

package poll

func SkipsCompletionPortOnSuccess(fd *FD) bool {
	return !fd.waitOnSuccess
}
