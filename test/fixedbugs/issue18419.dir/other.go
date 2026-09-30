package other

type Exported struct {
	Member int
}

func (e *Exported) member() int { return 1 }
