package main
/*
#include <errno.h>
typedef struct { int *p; int n; } c_buf;
static int add(c_buf *b, int n) { return (b ? b->n : 0) + n; }
static int divide(int a, int b) { if (!b) { errno = EDOM; return -1; } return a / b; }
*/
import "C"
import (
 "fmt"
 "strconv"
 "strings"
)
var trace []string
func mark(s string, n int) int { trace = append(trace, s); return n }
func must(ok bool) { if !ok { panic("failed assertion: " + strings.Join(trace, ",")) } }
func scenario(present bool) int {
 a, b := C.c_buf{n: 20}, C.c_buf{n: 10}
 var h *holder
 if present { h = &holder{buffer: &a} }
 var selected *C.c_buf
 if h != nil { selected = h.buffer }
 if selected == nil { selected = buf("fallback", &b) }
 n := int(C.add(selected, C.int(mark("tail", 2))))
 var f func() C.int
 if present { f = func() C.int { return 0 } }
 var call C.int
 if f != nil { call = f() }
 must(call == 0)
 var entries = map[int]*C.c_buf{}
 if present { entries[0] = &a }
 if entries[0] == nil { entries[0] = buf("init", &b) }
 must(entries[0] != nil)
 return n
}
type holder struct { buffer *C.c_buf }
func buf(s string, p *C.c_buf) *C.c_buf { trace = append(trace, s); return p }
func main() {
 for _, present := range []bool{false, true} {
  trace = nil
  n := scenario(present)
  want, calls := 12, "fallback,tail,init"
  if present { want, calls = 22, "tail" }
  must(n == want && strings.Join(trace, ",") == calls)
  fmt.Printf("present=%v n=%d trace=%s\n", present, n, strings.Join(trace, ","))
 }
 fmt.Println("PASS")
}
var _ = strconv.Atoi
