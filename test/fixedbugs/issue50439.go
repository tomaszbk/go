// compile

package p

var data []struct {
	F string `tag`
}

var V = ([]struct{ F string })(data)
