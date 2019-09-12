package client

//会引入我们GetFibonacci所在的本地包
import (
	"ch14/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	//t.Log(series.square(5))
}

//执行结果 重名的init方法被执行两次
//init1
//init2
//=== RUN   TestPackage
//--- PASS: TestPackage (0.00s)
//package_test.go:10: [1 1 2 3 5] <nil>
//PASS
