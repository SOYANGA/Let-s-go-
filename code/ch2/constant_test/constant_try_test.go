package constant_test

import "testing"

/*
枚举表示
 */
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday

)

/**
分别用二进制的三个位标识，三种状态
*/
const (
	Readable   = 1 << iota //可读
	Writable               //可写
	Executable             //可执行
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	b := 1 //0001
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
}
