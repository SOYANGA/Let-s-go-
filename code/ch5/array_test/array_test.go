package array_test

import "testing"

/*
数组的使用 声明与定义
*/
func TestArrayInit(t *testing.T) {
	//数组的声明
	var array [3]int

	//数组的声明并初始化
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}

	t.Log(array[0], array[1]) // 0 0
	t.Log(arr1, arr3)         //[1 2 3 4] [1 3 4 5]
}

/*
数组元素的遍历
*/
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}

	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	/*
		类似于for each 的遍历方式 idex是索引值 e为数组中的元素
		可以省略idex 使用 _,代替 idex, 担当返回值占位
	*/
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}

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
数组的切片
左闭右开
不填写则是数组完全拷贝
且不能填写负数
*/
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec1 := arr3[:3]
	arr3Sec2 := arr3[3:]
	arr3Sec3 := arr3[:]
	//arr3Sec3 := arr3[:-1]
	t.Log(arr3Sec1) //[1 2 3]
	t.Log(arr3Sec2) //[4 5]
	t.Log(arr3Sec3) //[1 2 3 4 5]
}
