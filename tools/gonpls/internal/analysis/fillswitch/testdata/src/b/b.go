package b

type TypeB int

const (
	TypeBOne TypeB = iota
	TypeBTwo
	TypeBThree
)

type ExportedInterface interface {
	isExportedInterface()
}

type notExportedType struct{}

func (notExportedType) isExportedInterface() {}
