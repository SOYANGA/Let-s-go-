package type_test

import (
	"math"
	"testing"
)

type MyInt int64

/**
测试隐式类型转换的约束
结论：不支持隐式类型转换
	别名也被允许隐式类型转换
如果想要转换则需要显示的类型强转，但是需要注意类型强转可能会发生截短，导致数据精度损失和值溢出的问题
*/
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
	t.Log(math.MaxInt64, math.MaxFloat64, math.MaxUint32)
}

/**
测试指针的使用及限制
go语言中支持指针 但是不支持指针的运算
*/
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	c := &aPtr
	//不支持指针的运算
	//aPtr+= 1    //. \type_test.go:29:6: invalid operation: aPtr += 1 (mismatched types *int and int)
	t.Log(a, aPtr, c)              //1  0xc00004e290 0xc00007c028
	t.Logf("%T %T %T", a, aPtr, c) //int  *int **int
}

/**
string在go中是数值类型 在初始化时会初始化为""字符串 而不是nil
 */
func TestString(t *testing.T) {
	var s string
	t.Log("*" +s + "*") //**
	t.Log(len(s)) //0

	//判断一个字符串是否为空在go中如下判断
	if s==""{
		//...
	}
	if len(s) == 0{
		//...
	}
}
