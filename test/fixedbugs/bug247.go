// run

package main

func main() {
	const (
		Delta = 100 * 1e6
		Count = 10
	)
	_ = int64(Delta * Count)
	var i interface{} = Count
	j := i.(int)
	if j != Count {
		println("j=", j)
		panic("fail")
	}
}
