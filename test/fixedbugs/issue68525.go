// run -gcflags='all=-N -l'

package main

import "os"

func main() {
	os.OpenFile(os.DevNull, os.O_WRONLY, 0)
}
