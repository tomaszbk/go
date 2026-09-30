package bufio

// Exported for testing only.
import (
	"unicode/utf8"
)

var IsSpace = isSpace

const DefaultBufSize = defaultBufSize

func (s *Scanner) MaxTokenSize(n int) {
	if n < utf8.UTFMax || n > 1e9 {
		panic("bad max token size")
	}
	if n < len(s.buf) {
		s.buf = make([]byte, n)
	}
	s.maxTokenSize = n
}

// ErrOrEOF is like Err, but returns EOF. Used to test a corner case.
func (s *Scanner) ErrOrEOF() error {
	return s.err
}
