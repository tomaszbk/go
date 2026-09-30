// run


package main

func main() {
	m := make(map[string][1000]byte)
	m["hi"] = [1000]byte{1}
	
	v := m["hi"]
	
	for k, vv := range m {
		if k != "hi" || string(v[:]) != string(vv[:]) {
			panic("bad iter")
		}
	}
}
