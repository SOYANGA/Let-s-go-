package map_ext

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2)) // 2 4 8
}

/*
利用map实现Set
1. 添加元素
2. 判断是否存在
3. 删除元素
4. 元素个数
*/
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}

	//添加元素
	mySet[1] = true
	n := 3

	//判断元素是否存在
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n) //3 is not existing
	}

	//元素个数
	mySet[3] = true
	t.Log(len(mySet)) //2

	//删除元素
	delete(mySet, 1)
	t.Log(len(mySet)) //1
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n) // 1 is not existing
	}
}
