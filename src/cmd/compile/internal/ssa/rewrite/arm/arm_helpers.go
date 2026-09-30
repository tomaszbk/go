package arm

// check if an immediate can be directly encoded into an ARM's instruction.
func isARMImmRot(v uint32) bool {
	for i := 0; i < 16; i++ {
		if v&^0xff == 0 {
			return true
		}
		v = v<<2 | v>>30
	}

	return false
}
