package main

import (
	"fmt"
	"internal/trace"
	"internal/trace/traceviewer"
	"slices"
)

// viewerFrames returns the frames of the stack of ev. The given frame slice is
// used to store the frames to reduce allocations.
func viewerFrames(stk trace.Stack) []trace.StackFrame {
	return slices.Collect(stk.Frames())
}

func viewerGState(state trace.GoState, inMarkAssist bool) traceviewer.GState {
	switch state {
	case trace.GoUndetermined:
		return traceviewer.GDead
	case trace.GoNotExist:
		return traceviewer.GDead
	case trace.GoRunnable:
		return traceviewer.GRunnable
	case trace.GoRunning:
		return traceviewer.GRunning
	case trace.GoWaiting:
		if inMarkAssist {
			return traceviewer.GWaitingGC
		}
		return traceviewer.GWaiting
	case trace.GoSyscall:
		// N.B. A goroutine in a syscall is considered "executing" (state.Executing() == true).
		return traceviewer.GRunning
	default:
		panic(fmt.Sprintf("unknown GoState: %s", state.String()))
	}
}
