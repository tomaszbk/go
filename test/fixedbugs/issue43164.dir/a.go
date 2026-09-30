package p

import . "strings"

var _ = Index // use strings

type t struct{ Index int }

var _ = t{Index: 0}
