// asmcheck


package codegen

var wsp = [256]bool{
	' ':  true,
	'\t': true,
	'\n': true,
	'\r': true,
}

func zeroExtArgByte(ch [2]byte) bool {
	return wsp[ch[0]] // amd64:-"MOVBLZX ..,.."
}

func zeroExtArgUint16(ch [2]uint16) bool {
	return wsp[ch[0]] // amd64:-"MOVWLZX ..,.."
}
