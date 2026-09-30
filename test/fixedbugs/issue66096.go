// compile

package p

type Message struct {
	Header map[string][]string
}

func f() {
	m := Message{Header: map[string][]string{}}
	m.Header[""] = append([]string(m.Header[""]), "")
	_ = m
}
