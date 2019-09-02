package operator_test

import "testing"

/**
分别用二进制的三个位标识，三种状态
*/
const (
	Readable   = 1 << iota //可读
	Writable               //可写
	Executable             //可执行
)

/**
数组的比较
- 相同维数且含有相同个数的元素的数组才可以进行比较
- 每个元素都相同的才相等，反之不等
*/
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c:=[...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) //false
	//t.Log(a==c) //.\operator_test.go:11:9: invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d) //true
}

/*
&^清零操作符的使用
*/
func TestBitClear(t *testing.T) {
	a := 7 //0111

	//清除读和可执行功能的功能
	a = a &^ Readable
	a &^= Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable) //false true false
}
