package open_close

import "fmt"

type Basic interface {
	Save()     //  存款业务
	Withdraw() //取款业务
	Transfer() //转账业务
}

type BasicBank struct{}

func (b *BasicBank) Save() {
	fmt.Println("basic  -- Save")
}

func (b *BasicBank) Withdraw() {
	fmt.Println("basic  -- Withdraw")
}

func (b *BasicBank) Transfer() {
	fmt.Println("basic  -- Transfer")
}

//增加一个方法，而不是去Basic interface增加方法，改用interface组合
func (b *BasicBank) Business() {
	fmt.Println("third party -- Business")
}

type Comprehensive interface {
	Basic
	Business()
}

func Demo() {
	b := &BasicBank{}
	b.Business()
}
