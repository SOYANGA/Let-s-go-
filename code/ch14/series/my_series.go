package series

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}



//大写字母开头包外可以访问
func GetFibonacci(n int) ([]int, error) {
	fibList := []int{1, 1}
	for i := 2; /*短变量声明*/ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

//小写字母开头包外不能访问
func square(n int) int {
	return n * n
}
