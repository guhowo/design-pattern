package decorator

import (
	"fmt"
	"reflect"
)

type IStream interface {
	Read()
	Write()
}

type FileStream struct {
	Name string
}

func (f *FileStream) Read() {
	fmt.Println("read file stream..." + f.Name)
}

func (f *FileStream) Write() {
	fmt.Println("write file stream...", f.Name)
}

type NetWorkStream struct{}

func (n *NetWorkStream) Read() {
	fmt.Println("read network stream...")
}

func (n *NetWorkStream) Write() {
	fmt.Println("write network stream...")
}

type MemoryStream struct{}

func (n *MemoryStream) Read() {
	fmt.Println("read memory stream...")
}

func (n *MemoryStream) Write() {
	fmt.Println("write memory stream...")
}

//定义Crypt装饰器类；不一定要定义装饰器接口
//type ICrypt interface {
//	Do()
//}

//crypt装饰器1，同样实现了stream接口。装饰器类和具体的类都实现了同一个基类
//如果这里用继承，即filestream struct，而不是istream，那么类的数量会爆炸
type CryptStream struct {
	Stream IStream
	Salt   string //加密用的信息
}

func (c *CryptStream) Do() {
	fmt.Println("crypt with slat ", c.Salt)
}

func (c *CryptStream) Read() {
	c.Stream.Read()
	c.Do()
}

func (c *CryptStream) Write() {
	c.Stream.Write()
}

//带buffer的装饰器2，同样实现了stream接口。装饰器类和具体的类都实现了同一个基类
type BufferedStream struct {
	Stream IStream
}

func (c *BufferedStream) Do() {
	fmt.Println("buffered ", reflect.TypeOf(c.Stream))
}

func (c *BufferedStream) Read() {
	c.Stream.Read()
	c.Do()
}

func (c *BufferedStream) Write() {
	c.Stream.Write()
}
