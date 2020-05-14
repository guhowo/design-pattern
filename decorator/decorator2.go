package decorator

import (
	"fmt"
	"reflect"
)

//区别于decorator.go文件中的实现方式，此处采用函数式编程的方式实现装饰器模式。

//装饰器1：加密装饰器
func Crypt(f func()) func(string) error {
	return func(salt string) error {
		f()
		fmt.Printf("crypt with %s...\n", salt)
		return nil
	}
}

//装饰器2：缓冲装饰器
func Buffer(f func()) func(reflect.Type) error {
	return func(p reflect.Type) error {
		f()
		fmt.Println("buffered:", p)
		return nil
	}
}
