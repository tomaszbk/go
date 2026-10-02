package gonflow

import "testing"

func check(t *testing.T, present bool, input string) {
 t.Helper()
 good := input != "bad"
 g := Global(input)
 if good && g != 7 || !good && g != -1 { t.Fatal(g) }
 n, err := Lambda(input)
 if good { if n != 8 || err != nil { t.Fatal(n, err) } } else if n != 0 || err == nil { t.Fatal(n, err) }
 var p *node
 var v *int
 if present { p = &node{}; v = value(9) }
 n, err = Chain(p, input)
 if !present { if n != -1 || err != nil { t.Fatal(n, err) } } else if good { if n != 7 || err != nil { t.Fatal(n, err) } } else if n != 0 || err == nil { t.Fatal(n, err) }
 for _, f := range []func(*int, string) (int, error){Coalesce, Assign} {
  n, err = f(v, input)
  if present { if n != 9 || err != nil { t.Fatal(n, err) } } else if good { if n != 7 || err != nil { t.Fatal(n, err) } } else if n != 0 || err == nil { t.Fatal(n, err) }
 }
}
func TestPresentOK(t *testing.T) { check(t, true, "7") }
func TestAbsentOK(t *testing.T) { check(t, false, "7") }
func TestPresentError(t *testing.T) { check(t, true, "bad") }
func TestAbsentError(t *testing.T) { check(t, false, "bad") }
