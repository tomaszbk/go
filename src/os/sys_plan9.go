package os

func hostname() (name string, err error) {
	f, err := Open("#c/sysname")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf [128]byte
	n, err := f.Read(buf[:len(buf)-1])

	if err != nil {
		return "", err
	}
	if n > 0 {
		buf[n] = 0
	}
	return string(buf[0:n]), nil
}
