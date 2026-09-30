// compile

package p

func New() resource {
	return &Client{}
}

type resource interface {
	table()
}

type Client struct {
	m map[Key1]int
}

func (c *Client) table() {}

type Key1 struct {
	K Key2
}

type Key2 struct {
	f [2]any
}
