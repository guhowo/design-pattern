package single

import "fmt"

type Basic interface {
	Save()     //  存款业务
	Withdraw() //取款业务
	Transfer() //转账业务
}

type Finance interface {
	Stock()   //股票
	Futures() //期货
	Fund()    //基金
}

type FinanceBasic interface {
	Finance
	Basic
}

//比如现在要开一个基本业务的网点，只需要实现三种方法
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

//此时业务扩张，这个网店要增加金融业务
type FinanceBank struct{}

func (f *FinanceBank) Finance() {
	fmt.Println("financial  -- Finance")
}

func (f *FinanceBank) Basic() {
	fmt.Println("financial  -- Basic")
}

type BankerPlus struct {
	BasicBank
	FinanceBank
}
type BankerT struct {
	Basic
}

func Demo() {
	bp := &BankerPlus{}
	bp.Save()
	bp.Withdraw()
	bp.Transfer()
	bp.Finance()
	bp.Basic()
}
