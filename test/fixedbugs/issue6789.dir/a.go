package a

type unexported struct {
        a int
        b bool
}

type Struct struct {
        unexported
}
