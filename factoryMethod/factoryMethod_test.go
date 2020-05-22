package factoryMethod

import "testing"

func Test_Factory(t *testing.T) {
	var factory ISplitterFactory
	factory = &BinarySplitterFactory{}
	splitter := factory.CreateSplitter()
	splitter.Split()

	factory = &TextSplitterFactory{}
	splitter = factory.CreateSplitter()
	splitter.Split()
}
