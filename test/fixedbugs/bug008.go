// run


package main

func main() {
	i5 := 5;

	switch {  // compiler crash fixable with 'switch true'
	case i5 < 5: dummy := 0; _ = dummy;
	case i5 == 5: dummy := 0; _ = dummy;
	case i5 > 5: dummy := 0; _ = dummy;
	}
}
/*
Segmentation fault
*/
