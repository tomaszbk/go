// Imported by private.go, which should not be able to see the private method.

package p

type Exported interface {
	private()
}

type Implementation struct{}

func (p *Implementation) private() {}

var X = new(Implementation)

