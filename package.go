package awesomelib

type ExportedType interface {
	PublicAction()
}

type internalType struct {
	internalNumber	int
}

func (a internalType) PublicAction() { /* I work in silence*/ }

func internalDityWork() { /* TMI*/ }

func NewAwesomeness() ExportedType { /* TMI*/ return internalType{internalNumber: 42} }
