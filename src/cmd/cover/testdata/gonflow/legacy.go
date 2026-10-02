package gonflow

func Lambda(s string) (int, error) {
 var f func() (int, error) = func() (int, error) { // @lambda-create
  v, err := parse(s) // @lambda-parse
  if err != nil { return 0, err }
  v++ // @lambda-increment
  return v, nil
 }
 v, err := f() // @lambda-call
 return v, err // @lambda-after
}

func Chain(p *node, s string) (int, error) {
 var v *int // @chain-select
 if p != nil {
  n, err := parse(s)
  if err != nil { return 0, err }
  v = p.Get(n)
 }
 if v == nil { v = value(-1) }
 n := *v // @chain-after
 return n, nil
}

func Coalesce(p *int, s string) (int, error) {
 v := p // @coalesce-select
 if v == nil {
  var err error
  v, err = load(s)
  if err != nil { return 0, err }
 }
 n := *v // @coalesce-after
 return n, nil
}

func Assign(p *int, s string) (int, error) {
 m := map[int]*int{0:p}
 if m[0] == nil { // @assign-select
  v, err := load(s)
  if err != nil { return 0, err }
  m[0] = v
 }
 n := *m[0] // @assign-after
 return n, nil
}

var global func(string) int = func(s string) int {
 v, err := parse(s)
 if err != nil {
  return -1 // @global-handler
 }
 return v
}
func Global(s string) int { return global(s) }
