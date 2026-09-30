package p

type m struct {
	link *m
}

var head *m

func F(m *int) bool {
	return head != nil
}
