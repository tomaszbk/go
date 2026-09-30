package io

type Writer interface {
	WrongWrite()
}

type SectionReader struct {
	X int
}

func SR(*SectionReader) {}
