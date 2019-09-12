// Go错误机制的练习
package err__test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

//区分错误类型——预制错误类型法（最常见的方法）
//定义不同的错误变量，以便于判断错误类型
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

//斐波那契额数列的动态规划实现(直接可以使用切片简单实现)
func GetFibonacci(n int) ([]int, error) {
	//对参数进行快速失败检测处理
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; /*短变量声明*/ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

//对斐波那契额数列的测试
//输入10
// 输出[1 1 2 3 5 8 13 21 34 55] <nil>
//发现的缺陷如果传入的n的参数是一个负数eg:-10则输出结果并不是我们想要[1 1] <nil>
//所以斐波那契额数列需要对传入的参数进行参数校验以及错误处理
func TestFibonacci(t *testing.T) {
	//[1 1 2 3 5 8 13 21 34 55] <nil>
	//t.Log(GetFibonacci(10))
	//
	//[] n should be in [2,100]
	//t.Log(GetFibonacci(-10))

	//错误检查+错误类型检查+错误恢复
	if v, err := GetFibonacci(-1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("Need a larger number")
			//...错误恢复处理
		}
		if err == LargerThanHundredError {
			fmt.Println("Need a less number")
			//...错误恢复处理
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

//传入字符串->转化为int然后求斐波那契额数列的值(错误版本 嵌套错误处理)
func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

//及早失败，避免嵌套
//传入字符串->转化为int然后求斐波那契额数列的值
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
