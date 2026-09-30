package p

type Exported interface {
	private()
}

type Implementation struct{}

func (p *Implementation) private() { println("p.Implementation.private()") }

var X = new(Implementation)
