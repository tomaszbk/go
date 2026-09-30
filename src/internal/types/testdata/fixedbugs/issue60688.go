package p

type String string

func g[P any](P, string) {}

// String and string are not identical and thus must not unify
// (they are element types of the func type and therefore must
// be identical to match).
// The result is an error from type inference, rather than an
// error from an assignment mismatch.
var f func(int, String) = g // ERROR "inferred type func(int, string) for func(P, string) does not match type func(int, String) of f"
