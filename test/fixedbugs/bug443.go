// compile


// Was failing to compile with 'invalid receiver' due to
// incomplete type definition evaluation.  Issue 3709.

package p

type T1 struct { F *T2 }
type T2 T1

type T3 T2
func (*T3) M()  // was invalid receiver

