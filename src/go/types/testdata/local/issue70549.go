package p

import "math"

var (
	_ = math.Sqrt
	_ = math.SQrt /* ERROR "undefined: math.SQrt (but have Sqrt)" */
	_ = math.sqrt /* ERROR "undefined: math.sqrt (but have Sqrt)" */
	_ = math.Foo  /* ERROR "undefined: math.Foo" */
	_ = math.foo  /* ERROR "undefined: math.foo" */
)
