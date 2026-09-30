package a

type myError string

func (e myError) Error() string { return string(e) }

const myErrorVal myError = "error"

func IsMyError(err error) bool {
	return err == error(myErrorVal)
}
