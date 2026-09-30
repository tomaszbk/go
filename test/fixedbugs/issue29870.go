// compile

// Make sure we can compile "_" functions without crashing.

package main

import "log"

func _() {
	log.Println("%2F")
}
