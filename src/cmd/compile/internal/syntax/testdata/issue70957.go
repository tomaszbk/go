package p

func f() { goto /* ERROR syntax error: unexpected semicolon, expected name */ ;}

func f() { goto } // ERROR syntax error: unexpected }, expected name
