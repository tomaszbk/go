package issue41761a

/*
   typedef struct S41761 S41761;
*/
import "C"

type T struct {
	X *C.S41761
}
