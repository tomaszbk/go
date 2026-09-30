// compile

package main

func f() map[string]interface{} {
	var p *map[string]map[string]interface{}
	_ = p
	return nil
}

func main() {
	f()
}
