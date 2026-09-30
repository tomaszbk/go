// build


package main

const big uint64 = 1<<63

func f(a uint64) uint64 {
	return a << big
}

func main() {
	f(1)
}

/*
main·f: doasm: notfound from=75 to=13 (82)    SHLQ    $-9223372036854775808,BX
main·f: doasm: notfound from=75 to=13 (82)    SHLQ    $-9223372036854775808,BX
main·f: doasm: notfound from=75 to=13 (82)    SHLQ    $-9223372036854775808,BX
*/
