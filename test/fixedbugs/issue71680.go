// compile

package p

type Parser struct{}
type Node struct{}

type parserState func(p *Parser) parserState

func parserStateData(root *Node) parserState {
	return func(p *Parser) parserState {
		return parserStateOpenMap(root)(p)
	}
}

func parserStateOpenMap(root *Node) parserState {
	return func(p *Parser) parserState {
		switch {
		case p != nil:
			return parserStateData(root)(p)
		}
		return parserStateOpenMap(root)(p)
	}
}
