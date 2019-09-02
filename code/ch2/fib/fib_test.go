package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//1.
	//var a int = 1
	//var b int = 1

	//2.
	//var(
	//	a int= 1
	//	b = 1
	//)

	//3.推荐 使用类型推断 直接给变量赋值
	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		//temp := a
		//a = b
		//b = temp + a
		a, b = b, a+b
	}
}

/**
交换两个变量的值
*/
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//temp := a
	//a = b
	//b = temp
	a, b = b, a
	t.Log(a, b)
}
