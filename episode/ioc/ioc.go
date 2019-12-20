package ioc

import "fmt"

//面向接口去编程，而不是面向具体的对象去编程！！！

type Car interface {
	Run()
}
type Driver interface {
	Drive(Car)
}

type Benz struct{}
type BMW struct{}

type James struct{}
type Sam struct{}

func (o *James) Drive(car Car) {
	fmt.Printf("Jame is driving ")
	car.Run()
}

func (o *Sam) Drive(car Car) {
	fmt.Printf("Sam is driving ")
	car.Run()
}

func (c *Benz) Run() {
	fmt.Println("Benz. Run")
}

func (c *BMW) Run() {
	fmt.Println("BMW. Run")
}

func Demo() {
	benz, bmw := &Benz{}, &BMW{}
	james, sam := &James{}, &Sam{}

	james.Drive(benz)
	james.Drive(bmw)
	sam.Drive(benz)
	sam.Drive(bmw)
}
