package pkg

func NonASCII(b []byte, i int) int {
	for i = 0; i < len(b); i++ {
		if b[i] >= 0x80 {
			break
		}
	}
	return i
}

