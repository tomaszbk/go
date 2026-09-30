package defers

import "time"

func _() {
	defer time.Since(time.Now()) // ERROR "call to time.Since is not deferred"
}
