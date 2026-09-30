package issue46404

// TODO(gri) re-enable this test with matching errors
//           between go/types and types2
// Check that we don't type check t[_] as an instantiation.
// type t [t /* type parameters must be named */ /* not a generic type */ [_]]_ // cannot use
