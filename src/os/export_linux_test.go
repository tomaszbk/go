package os

var (
	PollCopyFileRangeP  = &pollCopyFileRange
	PollSpliceFile      = &pollSplice
	GetPollFDAndNetwork = getPollFDAndNetwork
	CheckPidfdOnce      = checkPidfdOnce
)

const StatusDone = statusDone

func (p *Process) Status() processStatus {
	return processStatus(p.state.Load())
}
