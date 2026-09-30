// compile

package main

func main() {
	Get([]string{"a", "b"})
}

func Get(ss []string) *[2]string {
	return (*[2]string)(ss)
}
