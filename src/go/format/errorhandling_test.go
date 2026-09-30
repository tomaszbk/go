package format

import (
	"bytes"
	"go/parser"
	"go/token"
	"testing"
)

func TestErrorHandling(t *testing.T) {
	for _, src := range []string{
		"package p; func f() error { x:=read()!; _=x; return nil }",
		"package p; func f() error { x:=read() or err {return err}; _=x; return nil }",
		"package p; func f() error { return use(read()!) }",
		"package p; func f() error { if read() or err { return err } { }; return nil }",
		"package p; func f() error { read()! // trailing\nreturn nil }",
		"package p; func f() error { read()! /* trailing */\nreturn nil }",
	} {
		formatted, err := Source([]byte(src))
		if err != nil {
			t.Fatalf("format %q: %v", src, err)
		}
		if _, err := parser.ParseFile(token.NewFileSet(), "p.go", formatted, 0); err != nil {
			t.Fatalf("formatted source is invalid: %v\n%s", err, formatted)
		}
		again, err := Source(formatted)
		if err != nil || !bytes.Equal(formatted, again) {
			t.Errorf("format is not idempotent: %v\n%s\n%s", err, formatted, again)
		}
	}
}
