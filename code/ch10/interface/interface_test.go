package interface__test

import "testing"

/*
定义一个接口 interface
*/
type Programmer interface {
	WriteHelloWorld() string
}

/*
定义一个结构 AImpl
*/
type GoProgrammer struct {
}

/*
实现这个接口 ->该结构绑定的方法与接口的方法参数跟放回值完全相同
go中的接口是非侵入式，实现不依赖于接口定义，所以接口的定义可以包含在接口使用者包内
要实现这个接口 ->该结构绑定的方法与接口的方法参数跟放回值完全相同即可
*/
func (p *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

/*
AClient
*/
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld()) //fmt.Println("Hello World")

	var programmer Programmer = &GoProgrammer{} //fmt.Println("Hello World") 接口变量
	t.Log(programmer.WriteHelloWorld())
}
