// compile

// Issue 8079: gccgo crashes when compiling interface with blank type name.

package p

type _ interface{}
