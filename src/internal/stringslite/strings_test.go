package stringslite_test

import (
	"internal/stringslite"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestIsSpace(t *testing.T) {
	for r := rune(0); r <= utf8.MaxRune; r++ {
		if stringslite.IsSpace(r) != unicode.IsSpace(r) {
			t.Fatalf("IsSpace(%U) = %v, want %v", r, stringslite.IsSpace(r), unicode.IsSpace(r))
		}
	}
}
