// run


package main

func main() {
	var p *[3]int
	for i, _ := range *p {
		_ = i
	}
	for i, _ := range (*p) { // Note the parentheses! gofmt wants to remove them - don't let it!
		_ = i
	}
	var i int
	for i, (_) = range *p { // Note the parentheses! gofmt wants to remove them - don't let it!
		_ = i
	}
}
