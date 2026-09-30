// run

package main

func main() {
	data := make(map[int]string, 1)
	data[0] = "hello, "
	data[0] += "world!"
	if data[0] != "hello, world!" {
		panic("BUG: " + data[0])
	}
}
