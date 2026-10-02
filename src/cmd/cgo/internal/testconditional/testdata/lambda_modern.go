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
func parse(s string) (int, error) {
 mark("create", 0)
 var fn func(string) (int, error) = (s) => {
  mark("call", 0)
  n := strconv.Atoi(s)!
  return n + 1, nil
 }
 n, err := invoke(fn, s)
 mark("after", 0)
 return n, err
}
func cparse(d int) (int, error) {
 var fn func() (int, error) = () => {
  n := C.divide(9, C.int(d))!
  return int(n), nil
 }
 return fn()
}
func pointerArg(b *C.c_buf) int {
 return int(C.add(b, value(() => C.int(strconv.Atoi("2") or err { return 0 }))))
}
func invoke(f func(string) (int, error), s string) (int, error) { return f(s) }
func value(f func() C.int) C.int { return f() }
func main() {
 for _, input := range []string{"6", "bad"} {
  trace = nil
  n, err := parse(input)
  must((input == "6" && n == 7 && err == nil) || (input == "bad" && n == 0 && err != nil))
  must(strings.Join(trace, ",") == "create,call,after")
  fmt.Printf("parse %s n=%d error=%v trace=%s\n", input, n, err != nil, strings.Join(trace, ","))
 }
 for _, divisor := range []int{3, 0} {
  n, err := cparse(divisor)
  must((divisor == 3 && n == 3 && err == nil) || (divisor == 0 && n == 0 && err != nil))
  fmt.Printf("C divide %d n=%d error=%v\n", divisor, n, err != nil)
 }
 var b C.c_buf
 b.n = 10
 must(pointerArg(&b) == 12)
 fmt.Println("PASS")
}
