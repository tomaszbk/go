package p

import "io"

type A interface {
        io.Reader
}

func f(a A) {
        a.Read(nil)
}
