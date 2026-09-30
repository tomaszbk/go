// This file contains tests for the unmarshal checker.

package unmarshal

import "encoding/json"

func _() {
	type t struct {
		a int
	}
	var v t

	json.Unmarshal([]byte{}, v) // ERROR "call of Unmarshal passes non-pointer as second argument"
}
