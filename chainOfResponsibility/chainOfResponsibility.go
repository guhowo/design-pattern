package chainOfResponsibility

import "fmt"

// 定义处理器类的抽象父类
type IHandler interface {
	handle(*string) bool
}

// 实现处理器类的接口方法 A
type PlusA struct{}

func (a *PlusA) handle(content *string) bool {
	fmt.Println(*content)
	*content = *content + " aaa"
	return true
}

// 实现处理器类的接口方法 A
type PlusB struct{}

func (b *PlusB) handle(content *string) bool {
	fmt.Println(*content)
	*content = *content + " bbb"
	return true
}

// 定义职责链的接口
type IHandlerChain interface {
	Init()
	AddHandler(handler IHandler)
	DoHandle() bool
}

type HandlerChain struct {
	content  *string
	handlers []IHandler //存储职责链上的对象
}

func (c *HandlerChain) Init(content string) {
	c.handlers = make([]IHandler, 0)
	c.content = &content
}

func (c *HandlerChain) AddHandler(handler IHandler) {
	c.handlers = append(c.handlers, handler)
}

func (c *HandlerChain) DoHandle() bool {
	for _, handler := range c.handlers {
		if !handler.handle(c.content) {
			return false
		}
	}
	fmt.Println(*c.content)
	return true
}
