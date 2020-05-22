package factoryMethod

import "fmt"

// 工厂方法的接口
type ISplitterFactory interface {
	CreateSplitter() ISplitter
}

// 文件分割器接口
type ISplitter interface {
	Split()
}

// concrete splitter:1
type BinarySplitter struct{}

func (s *BinarySplitter) Split() {
	fmt.Println("split binary")
}

// concrete splitter:2
type TextSplitter struct{}

func (s *TextSplitter) Split() {
	fmt.Println("text binary")
}

//每个类都有一个对应的工厂类
type BinarySplitterFactory struct{}

func (c *BinarySplitterFactory) CreateSplitter() ISplitter {
	return &BinarySplitter{}
}

type TextSplitterFactory struct{}

func (c *TextSplitterFactory) CreateSplitter() ISplitter {
	return &TextSplitter{}
}
