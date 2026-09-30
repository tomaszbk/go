// compile

package main

import "fmt"

func main() {
	var ii interface{} = 5
	zz, err := ii.(interface{})
	fmt.Println(zz, err)
}
