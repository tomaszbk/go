package runtime

func XTestSPWrite(t TestingT) {
	// Test that we can traceback from the stack check prologue of a function
	// that writes to SP. See #62326.

	// Start a goroutine to minimize the initial stack and ensure we grow the stack.
	done := make(chan bool)
	go func() {
		testSPWrite() // Defined in assembly
		done <- true
	}()
	<-done
}
