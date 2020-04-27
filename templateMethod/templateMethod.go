package templateMethod

import "fmt"

/*
	模板方法适合在主要流程很稳定或者不变化的情况下，
	具体的某些子步骤需要改动的情况。
	模板方法一般实现为：定义一个父类（实现基本不变的服务），变化的服务设计为虚函数，让子类实现。
*/

//接口
type Cooker interface {
	fire()
	cooke()
}

//父类
type Cook struct{}

func (c *Cook) fire() {
	fmt.Println("CookMenu fire")
}

func (c *Cook) cooke() {
	fmt.Println("CookMenu cooke")
}

//子类1：可以选择性的实现某些方法
type CookTomato struct {
	Cook
}

func (c *CookTomato) cooke() {
	fmt.Println("Cook Tomato")
}

//子类1：可以选择性的实现某些方法
type CookOlive struct {
	Cook
}

func (c *CookOlive) fire() {
	fmt.Println("olive fresh")
}

//主框架稳定
func doCook(cooker Cooker) {
	cooker.fire()
	cooker.cooke()
}
