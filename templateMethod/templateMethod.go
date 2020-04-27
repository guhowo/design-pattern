package templateMethod

import "fmt"

type Cooker interface {
	fire()
	cooke()
}

type Cook struct{}

func (c *Cook) fire() {
	fmt.Println("CookMenu fire")
}

func (c *Cook) cooke() {
	fmt.Println("CookMenu cooke")
}

type CookTomato struct {
	Cook
}

func (c *CookTomato) cooke() {
	fmt.Println("Cook Tomato")
}

//主框架稳定
func doCook(cooker Cooker) {
	cooker.fire()
	cooker.cooke()
}
