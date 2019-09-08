package customer_type

import (
	"fmt"
	"testing"
	"time"
)

/*
相当于一个类型别名
*/
type IntConventionFn func(op int) int

/*
计算函数时常 —— 函数式编程 （函数可以作为参数，返回值，变量）
参数是函数类型 返回也是函数类型
参数函数只负责执行，返回的函数类型做了一层包装进行计算所用消耗的时间——类似于装饰着模式
不需要在每个函数侵入式的写入计算时间的代码，只需要调用这个方法传入函数则就可以计算处函数执行所消耗的时间

我们拆解一下这个函数 方法参数是一个函数 返回值是一个函数 显得方法很长不便于阅读，则我们自定类型来缩减一下这个方法
*/
func timeSpent(inner IntConventionFn) IntConventionFn {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)                                         //参数中传入的函数
		fmt.Println("time spent:", time.Since(start).Seconds()) // slowFun time spent: 1.0003273
		return ret
	}
}

/*
单纯的sleep1秒
*/
func slowFun(op int) int {
	time.Sleep(time.Second) //睡一秒
	return op
}

/*
测试函数作为变量和参数返回值——函数式编程 非侵入式
*/
func TestTimeSpent(t *testing.T) {
	tsSF := timeSpent(slowFun)       //函数可以作为变量
	t.Log("slowFun coast", tsSF(10)) //调用函数
}
