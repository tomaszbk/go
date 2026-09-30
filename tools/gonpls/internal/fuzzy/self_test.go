package fuzzy_test

import (
	"testing"

	. "golang.org/x/tools/gopls/internal/fuzzy"
)

func BenchmarkSelf_Matcher(b *testing.B) {
	idents := collectIdentifiers(b)
	patterns := generatePatterns()

	for b.Loop() {
		for _, pattern := range patterns {
			sm := NewMatcher(pattern)
			for _, ident := range idents {
				_ = sm.Score(ident)
			}
		}
	}
}

func BenchmarkSelf_SymbolMatcher(b *testing.B) {
	idents := collectIdentifiers(b)
	patterns := generatePatterns()

	for b.Loop() {
		for _, pattern := range patterns {
			sm := NewSymbolMatcher(pattern)
			for _, ident := range idents {
				_, _ = sm.Match([]string{ident})
			}
		}
	}
}
