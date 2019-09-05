package string

import (
	"testing"
)

/*
字符串与编码的区别
*/
func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值 ""
	s = "hello"
	t.Log(len(s)) //5 byte数
	//s[1] = '3' //string是不可变的byte slice //cannot assign to s[1]
	s = "\xE4\xB8\xA5" //可以存放任何二进制数据 16进制串
	//s = "\xE4\xB8\xFF"
	t.Log(s)      //严
	t.Log(len(s)) //是byte数 //3

	s = "中"
	t.Log(len(s)) //是byte数 //3

	c := []rune(s) //可以取出字符串的Unicode go语言的内置机制 字符串转化为rune的切片Unicode
	t.Log(len(c))  //1 unicode的编码的字符

	//t.Log("rune size:", unsafe.Sizeof(c[0])) //4

	t.Logf("中 unicode %x", c[0]) //中 unicode 4e2d 16进制 编码的值
	t.Logf("中 UTF8 %x", s)       //中 UTF8 e4 b8 ad 16进制 3byte
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	//迭代range输出的是rune不是byte 而单纯的for循环输出的是UTF8编码的字符
	for _, c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c) //[1]代表都是和第一个参数c对应 只是同一个参数的格式化方式不同
	}
}
//string_test.go:36: 中 20013 4e2d
//string_test.go:36: 华 21326 534e
//string_test.go:36: 人 20154 4eba
//string_test.go:36: 名 21517 540d
//string_test.go:36: 共 20849 5171
//string_test.go:36: 和 21644 548c
//string_test.go:36: 国 22269 56fd