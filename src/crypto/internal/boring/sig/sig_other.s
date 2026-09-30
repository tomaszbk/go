// These functions are no-ops.
// On amd64 they have recognizable implementations, so that you can
// search a particular binary to see if they are present.
// On other platforms (those using this source file), they don't.

//go:build !amd64

TEXT ·BoringCrypto(SB),$0
	RET

TEXT ·FIPSOnly(SB),$0
	RET

TEXT ·StandardCrypto(SB),$0
	RET
