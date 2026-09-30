package scannererr

import (
	"bufio"
	"io"
)

func _(r io.Reader) {
	s := bufio.NewScanner(r) // ERROR `bufio.Scanner .* is used in Scan loop at line 10 without final check of s.Err\(\)`
	for s.Scan() {
	}
}
