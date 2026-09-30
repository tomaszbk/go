// run


package main

type Type interface {
	TypeName() string;
}

type TInt struct {
}

// TInt
func (i *TInt) TypeName() string {
	return "int";
}


func main() {
	var t Type;
	t = nil;
	_ = t;
}

/*
bug023.go:20: fatal error: naddr: const <Type>I{<TypeName>110(<_t117>{},<_o119>{},{});}
*/
