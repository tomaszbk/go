package main

import (
	"errors"
	"fmt"
)

func must(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}

var trace string
var one, two = 1, 2
var sentinel = errors.New("failure")

type Node struct {
	Next  *Node
	Count int
	P     *int
	Items []*int
	Value any
	Fn    func(int) *int
}

func (n *Node) Get(x int) *int { trace += "m"; return n.P }
func (n *Node) Load(fail bool) (*int, error) {
	trace += "l"
	if fail {
		return &one, sentinel
	}
	return n.P, nil
}
func (n *Node) NilMethod() *int {
	trace += "n"
	if n == nil {
		return &two
	}
	return n.P
}
func (n Node) ByValue() *int { return n.P }

type Reader interface{ NilMethod() *int }
type Base struct{ P *int }
type Embedded struct{ *Base }
type MyErr struct{}

func (*MyErr) Error() string { return "error" }

type Errors struct{ E *MyErr }

func root(n *Node) *Node                { trace += "r"; return n }
func arg() int                          { trace += "a"; return 3 }
func fallback() *int                    { trace += "f"; return &two }
func first() int                        { trace += "1"; return 0 }
func last() int                         { trace += "2"; return 0 }
func collect(a int, p *int, b int) *int { return p }
func panics(f func()) (did bool)        { defer func() { did = recover() != nil }(); f(); return }
func main()                             { scenarios(); fmt.Println("PASS") }

func (n *Node) Clear()                   { trace += "z" }
func (n *Node) Multi(x int) (int, error) { trace += "q"; return x, nil }
