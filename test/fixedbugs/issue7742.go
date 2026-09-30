// compile


// Issue 7742: cannot use &autotmp_0001 (type *map[string]string) as type *string in function argument

package main

var (
	m map[string]string
	v string
)

func main() {
	m[v], _ = v, v
}
