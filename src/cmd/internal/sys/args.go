package sys

// ExecArgLengthLimit is the number of bytes we can safely
// pass as arguments to an exec.Command.
//
// Windows has a limit of 32 KB. To be conservative and not worry about whether
// that includes spaces or not, just use 30 KB. Darwin's limit is less clear.
// The OS claims 256KB, but we've seen failures with arglen as small as 50KB.
const ExecArgLengthLimit = (30 << 10)
