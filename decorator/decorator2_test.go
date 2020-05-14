package decorator

import (
	"reflect"
	"testing"
)

func TestDecorator(t *testing.T) {
	fs := FileStream{Name: "decorator"}
	Crypt(fs.Read)("fuck")

	Buffer(fs.Read)(reflect.TypeOf(fs))
}
