package string

import (
	"strconv"
	"strings"
	"testing"
)

/*
字符串的常用方法
1.字符串分割
2.字符串连接
*/
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Logf(part)
	}
	t.Log(strings.Join(parts, "-")) //A-B-C
}

/*
字符串与其他类型的转换 int <-> string
*/
func TestStrconv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Logf("str" + s) //str10
	if i, error := strconv.Atoi("10"); error == nil {
		t.Log(10 + i) //20
	}
}
