package decorator

import (
	"fmt"
	"reflect"
)

//区别于decorator.go文件中的实现方式，此处采用函数式编程的方式实现装饰器模式。

//装饰器1：加密装饰器
func Crypt(f func()) func(string) {
	return func(salt string) {
		f()
		fmt.Printf("crypt with %s...\n", salt)
	}
}

//装饰器2：缓冲装饰器
func Buffer(f func()) func(reflect.Type) {
	return func(p reflect.Type) {
		f()
		fmt.Println("buffered:", p)
	}
}
