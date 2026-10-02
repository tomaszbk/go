package gonflow

func Lambda(s string) (int, error) {
 var f func() (int, error) = () => { // @lambda-create
  v := parse(s)! // @lambda-parse
  v++ // @lambda-increment
  return v, nil
 }
 v, err := f() // @lambda-call
 return v, err // @lambda-after
}

func Chain(p *node, s string) (int, error) {
 v := p?.Get(parse(s)!) ?? value(-1) // @chain-select
 n := *v // @chain-after
 return n, nil
}

func Coalesce(p *int, s string) (int, error) {
 v := p ?? load(s)! // @coalesce-select
 n := *v // @coalesce-after
 return n, nil
}

func Assign(p *int, s string) (int, error) {
 m := map[int]*int{0:p}
 m[0] ??= load(s)! // @assign-select
 n := *m[0] // @assign-after
 return n, nil
}

var global func(string) int = (s) => parse(s) or err {
 return -1 // @global-handler
}
func Global(s string) int { return global(s) }
