// run


package main

import "os"

func main() {
	m := make(map[int]int);
	m[0] = 0;
	m[0]++;
	if m[0] != 1 {
		print("map does not increment\n");
		os.Exit(1)
	}
}
