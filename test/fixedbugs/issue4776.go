// errorcheck


// Issue 4776: missing package declaration error should be fatal.

type MyInt int32 // ERROR "package statement must be first|package clause"

