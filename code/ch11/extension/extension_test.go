package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Print(" ", host)
}

/*
Dog跟Pet组合
*/
type Dog struct {
	//复合
	//p *Pet
	//匿名组合
	Pet
}

////方法重载
//func (d *Dog) Speak() {
//	//d.p.Speak() //调用Pet的方法
//	fmt.Print("Wang!")
//}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Print(" ", host)
//}

func (d *Dog) Speak() {
	fmt.Print("Wang~")
}

/*
go无法支持LSP
*/
func TestDog(t *testing.T) {
	//dog := new(Dog)
	var dog = new(Dog)  // cannot use new(Dog) (type *Dog) as type Pet in assignment go中不支持显示类型转换的，不支持继承。则无法进行类型转换
	dog.SpeakTo("CHao") //内嵌的结构类型组合不支持LSP,不支持重载
	//dog.Speak() //wang~
}
