package b

import "issue30862.dir/a"

type EmbedImported struct {
	a.NoitfStruct
}

func Test() []string {
	bad := []string{}
	x := interface{}(new(a.NoitfStruct))
	if _, ok := x.(interface {
		NoInterfaceMethod()
	}); ok {
		bad = append(bad, "fail 1")
	}

	x = interface{}(new(EmbedImported))
	if _, ok := x.(interface {
		NoInterfaceMethod()
	}); ok {
		bad = append(bad, "fail 2")
	}
	return bad
}
