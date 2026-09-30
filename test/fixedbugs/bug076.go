// build


package main

func f() {
exit:
	;
	goto exit
}


func main() {
exit:
	; // this should be legal (labels not properly scoped?)
	goto exit
}

/*
uetli:~/Source/go/test/bugs gri$ 6g bug076.go 
bug076.go:11: label redeclared: exit
*/
