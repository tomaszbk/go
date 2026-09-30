// run


package main

import (
	"strings"
)

func main() {
	str1 := "你好世界"
	if !strings.HasSuffix(str1, "世界") {
		panic(1)
	}

	str2 := "こんにちは"
	if !strings.HasSuffix(str2, "ちは") {
		panic(2)
	}

	str3 := "спасибо"
	if !strings.HasSuffix(str3, "ибо") {
		panic(3)
	}
}
