package p

type (
        C[T any] interface{~int; M() T}

        _ C[bool]
        _ comparable
        _ interface {~[]byte | ~string}

        // Alias type declarations may refer to "constraint" types
        // like ordinary type declarations.
        _ = C[bool]
        _ = comparable
        _ = interface {~[]byte | ~string}
)
