package flate

// shiftMask is a no-op shift mask for x86-64.
// Using it lets the compiler omit the check for shift size >= 64.
const reg8SizeMask64 = 63
