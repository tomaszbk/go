// This file contains the test for canonical struct tags.

package structtag

type StructTagTest struct {
	A int "hello" // ERROR "`hello` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair"
}
