// errorcheck


package main

type	T struct
{
	f int;
}

// 6g used to get confused by the f:1 above
// and allow uses of f that would be silently
// dropped during the compilation.
var _ = f;	// ERROR "undefined"

var _ = T{f: 1}

