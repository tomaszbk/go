package main

import "strings"

func init() {
	register("stringconcat", stringconcat)
}

func stringconcat() {
	s0 := strings.Repeat("0", 1<<10)
	s1 := strings.Repeat("1", 1<<10)
	s2 := strings.Repeat("2", 1<<10)
	s3 := strings.Repeat("3", 1<<10)
	s := s0 + s1 + s2 + s3
	panic(s)
}
