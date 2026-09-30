package a

func Unique[T comparable](set []T) []T {
	nset := make([]T, 0, 8)

loop:
	for _, s := range set {
		for _, e := range nset {
			if s == e {
				continue loop
			}
		}

		nset = append(nset, s)
	}

	return nset
}
