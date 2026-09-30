// errorcheck


package f

func f(x int /* // GC_ERROR "unexpected newline"

*/) // GCCGO_ERROR "expected .*\).*|expected declaration"
