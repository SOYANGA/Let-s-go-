package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
函数多返回值测试
*/
func returnMulValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/*
使用_可以表示占位
*/
func TestReturnMulValues(t *testing.T) {
	a, b := returnMulValues()
	t.Log(a, b)
	_, c := returnMulValues()
	t.Log(c)
}

/*
计算函数时常 —— 函数式编程 （函数可以作为参数，返回值，变量）
参数是函数类型 返回也是函数类型
参数函数只负责执行，返回的函数类型做了一层包装进行计算所用消耗的时间——类似于装饰着模式
不需要在每个函数侵入式的写入计算时间的代码，只需要调用这个方法传入函数则就可以计算处函数执行所消耗的时间
*/
func timeSpent(inner func(op int) int) func(op int) int {
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

/*
函数求和参数
*/
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

/*
函数的可变参数
*/
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))    //10
	t.Log(Sum(1, 2, 3, 4, 5)) //15
}

/*
释放资源
 */
func clear() {
	fmt.Println("clear resource")
}

/*
测试defer函数
 */
func TestDefer(t *testing.T) {
	defer clear() //释放资源释放锁
	fmt.Println("Start")
	panic("err") //异常出错
	fmt.Println("End")
}
//输出
//Start
//clear resource

//Start
//clear resource
