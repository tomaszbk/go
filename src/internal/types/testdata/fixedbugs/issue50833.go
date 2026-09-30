package p

type (
	S  struct{ f int }
	PS *S
)

func a() []*S { return []*S{{f: 1}} }
func b() []PS { return []PS{{f: 1}} }

func c[P *S]() []P { return []P{{f: 1}} }
func d[P PS]() []P { return []P{{f: 1}} }
