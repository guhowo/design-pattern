package chainOfResponsibility

import "testing"

func TestHandlerChain_DoHandle(t *testing.T) {
	aHandler := &PlusA{}
	bHandler := &PlusB{}
	hc := HandlerChain{}

	hc.Init("hello, chain of responsibility. ")
	hc.AddHandler(aHandler)
	hc.AddHandler(bHandler)
	hc.DoHandle()
}
