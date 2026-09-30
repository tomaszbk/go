package cgroup

type LineReader = lineReader

func (l *LineReader) Next() error {
	return l.next()
}

func (l *LineReader) Line() []byte {
	return l.line()
}

func NewLineReader(fd int, scratch []byte, read func(fd int, b []byte) (int, uintptr)) *LineReader {
	return newLineReader(fd, scratch, read)
}

var (
	ErrEOF            = errEOF
	ErrIncompleteLine = errIncompleteLine
	ErrMalformedFile  = errMalformedFile
)

var ContainsCPU = containsCPU

var ParseV1Number = parseV1Number
var ParseV2Limit = parseV2Limit

var ParseCPUCgroup = parseCPUCgroup
var ParseCPUMount = parseCPUMount

var UnescapePath = unescapePath
