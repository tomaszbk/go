// run


// issue 10253: cmd/7g: bad codegen, probably regopt related

package main

func main() {
	if !eq() {
		panic("wrong value")
	}
}

var text = "abc"
var s = &str{text}

func eq() bool {
	return text[0] == s.text[0]
}

type str struct {
	text string
}
