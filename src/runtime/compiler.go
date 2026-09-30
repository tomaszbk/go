package runtime

// Compiler is the name of the compiler toolchain that built the
// running binary. Known toolchains are:
//
//	gc      Also known as cmd/compile.
//	gccgo   The gccgo front end, part of the GCC compiler suite.
const Compiler = "gc"
