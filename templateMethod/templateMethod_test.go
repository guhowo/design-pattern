package templateMethod

import "testing"

func TestDoCook(t *testing.T) {
	tomato := &CookTomato{}
	doCook(tomato)

	println()

	olive := &CookOlive{}
	doCook(olive)
}
