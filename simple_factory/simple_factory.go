package simple_factory

import "fmt"

type Operation struct {
	A int
	B int
}

type IOperation interface {
	GetResult() int
	SetA(a int)
	SetB(b int)
}

func (O *Operation) SetA(a int) {
	O.A = a
}

func (O *Operation) SetB(b int) {
	O.B = b
}

type AddOperation struct {
	Operation
}

func (this *AddOperation) GetResult() int {
	return this.A + this.B
}

type MulOperation struct {
	Operation
}

func (this *MulOperation) GetResult() int {
	return this.A * this.B
}

type IFactory interface {
	CreateOperation() Operation
}

type AddFactory struct{}

func (this *AddFactory) CreateOperation() IOperation {
	return &(AddOperation{})
}

func Try() {
	fac := &(AddFactory{})
	oper := fac.CreateOperation()
	oper.SetA(1)
	oper.SetB(2)
	fmt.Println(oper.GetResult())
}
