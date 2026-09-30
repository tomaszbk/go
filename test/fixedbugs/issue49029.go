// compile


package p

type s struct {
	f func()
}

func f() {
	ch := make(chan struct{}, 1)
	_ = [...]struct{ slice []s }{
		{}, {}, {}, {},
		{
			slice: []s{
				{
					f: func() { ch <- struct{}{} },
				},
			},
		},
	}
}
