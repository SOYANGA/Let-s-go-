package interface2__test

import (
	"fmt"
	"testing"
)

type Code string

/*
定义一个接口 interface
*/
type Programmer interface {
	WriteHelloWorld() Code
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
func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

/*
定义一个结构 AImpl
*/
type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

/*
公共方法
*/
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld()) //输出类型 以及调用方法返回的结果
}

/*
AClient
*/
func TestPolymorphism(t *testing.T) {
	goProgram := new(GoProgrammer)
	goProgram2 := &GoProgrammer{} //
	javaProgram := new(JavaProgrammer)
	writeFirstProgram(goProgram)   //*interface2__test.GoProgrammer fmt.Println("Hello World")
	writeFirstProgram(javaProgram) //*interface2__test.JavaProgrammer System.out.Println("Hello World")
	writeFirstProgram(goProgram2)  //*interface2__test.GoProgrammer fmt.Println("Hello World")
	fmt.Printf("%T\n",GoProgrammer{}) //interface2__test.GoProgrammer
	fmt.Printf("%T",&GoProgrammer{}) //*interface2__test.GoProgrammer
}
