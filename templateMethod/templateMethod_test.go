package templateMethod

import "testing"

func TestDoCook(t *testing.T) {
	tmt := &CookTomato{}
	doCook(tmt)
}
