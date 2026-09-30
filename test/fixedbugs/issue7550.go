// run


package main

func shouldPanic(f func()) {
        defer func() {
                if recover() == nil {
                        panic("not panicking")
                }
        }()
        f()
}

func f() {
        length := int(^uint(0) >> 1)
        a := make([]struct{}, length)
        b := make([]struct{}, length)
        _ = append(a, b...)
}

func main() {
	shouldPanic(f)
}
