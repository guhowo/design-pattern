package decorator

import "testing"

func TestCryptStream_Read(t *testing.T) {
	//文件加密
	fs := &FileStream{Name: "hellofile"}
	cryptDecorator := CryptStream{Stream: fs}
	cryptDecorator.Read()

	//network加密
	ns := &NetWorkStream{}
	cryptDecorator = CryptStream{Stream: ns}
	cryptDecorator.Read()

	//文件缓冲
	bufferedDecorator := BufferedStream{Stream: fs}
	bufferedDecorator.Read()

	//network缓冲
	bufferedDecorator = BufferedStream{Stream: ns}
	bufferedDecorator.Read()
}
