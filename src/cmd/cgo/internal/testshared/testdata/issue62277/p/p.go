package p

var S = func() []string {
	return []string{"LD_LIBRARY_PATH"}
}()

var T []string

func init() {
	T = func() []string {
		return []string{"LD_LIBRARY_PATH"}
	}()
}
