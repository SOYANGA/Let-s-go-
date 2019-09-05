package my_map

import "testing"

/*
map的声明及初始化
*/
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])                   //4
	t.Logf("len m1 = %d", len(m1)) //len m1 = 3

	m2 := map[int]int{}            //初始化为空的map
	m2[4] = 16                     //设定其中某个元素
	t.Logf("len m2 = %d", len(m2)) //len m2 = 1

	m3 := make(map[int]int, 2)
	t.Logf("len m3 = %d", len(m3)) //len m3 = 0 cap()不能用于求map的容量
}

/*
判断map中查询的值是否存在
*/
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) //0

	m1[2] = 0
	t.Log(m1[2]) //0

	//m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v) //key 3's value is 0
	} else {
		t.Log("Key 3 is not existing")
	}

}

/*
Map遍历 range 返回的两个值分别是key和value
*/
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

//map_test.go:46: 2 4
//map_test.go:46: 3 9
//map_test.go:46: 1 1

/*
Map元素的删除
*/
func TestDeleteMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	delete(m1, 1)
	t.Log(m1)       //map[2:4 3:9]
	//delete(nil, "1") //first argument to delete must be map; have nil
}
