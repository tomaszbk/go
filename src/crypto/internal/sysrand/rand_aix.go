package sysrand

func read(b []byte) error {
	return urandomRead(b)
}
