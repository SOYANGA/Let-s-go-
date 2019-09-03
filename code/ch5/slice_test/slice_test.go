package slice_test

import "testing"

/*
切片的初始化
*/
func TestSliceInit(t *testing.T) {

	//声明一个slice切片
	var s0 []int
	t.Log(len(s0), cap(s0)) //0 0

	//填充一个元素
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) //1 1

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1)) //4 4

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))    //3 5
	t.Log(s2[0], s2[1], s2[2]) // 0 0 0

	s2 = append(s2, 1)                //填充一个元素
	t.Log(s2[0], s2[1], s2[2], s2[3]) //0 0 0 1
	t.Log(len(s2), cap(s2))           //输出长度跟容量  4 5
}

/*
切片的扩容机制,扩容过程中切片中的指针可能会改变。
如果切片后面没有连续的存储空间，
则就会重新找一块可以容下连续存储空间进行将内容拷贝到其中
并改变切面中的指针

切面不像数组需要考虑固定的长度，可以动态增长，但是这个动态增长也会带来性能上的影响
*/
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

/*
   slice_test.go:37: 1 1
   slice_test.go:37: 2 2
   slice_test.go:37: 3 4
   slice_test.go:37: 4 4
   slice_test.go:37: 5 8
   slice_test.go:37: 6 8
   slice_test.go:37: 7 8
   slice_test.go:37: 8 8
   slice_test.go:37: 9 16
   slice_test.go:37: 10 16
每一次当cap不够时cap扩容时增长原来的2倍的cap
*/

/*
切片共享存储结构
*/
func TestSliceShareMemory(t *testing.T) {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := months[3:6]           //[Apr May Jun]
	t.Log(Q2, len(Q2), cap(Q2)) //3 9 //因为共享切面所以cap是9 (在months中Apr后面的连续存储空间)
	summer := months[5:8]
	t.Log(summer, len(summer), cap(summer)) //[Jun Jul Aug] 3 7 //跟Q2共享了一部分存储空间

	summer[0] = "Unknown"
	t.Log(Q2)     //[Apr May Unknown]
	t.Log(months) //[Jan Feb Mar Apr May Unknown Jul Aug Sep Oct Nov Dec]
}

/*
判断切片是否可以比较
*/
func TestSliceCompare(t *testing.T) {
	//a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	//if a == b { //invalid operation: a == b (slice can only be compared to nil)
	//	t.Log("equal")
	//}
}

/*
切片的复制 copy
如果参与复制的两个切片不一样大，则就会按照较小的那个切片的元素个数进行len
*/
func TestSliceCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	//i := copy(slice2, slice1)
	//t.Log(slice2,i) // [1 2 3] 3 只会复制slice1的前3个元素到slice2

	j := copy(slice1, slice2)
	t.Log(slice1, j) // [6 7 8 4 5] 3 只会复制slice2的3个元素到slice1的前3个元素

	slice3 := make([]int, 3, 5)
	slice4 := make([]int, 2, 5)
	slice3 = append(slice3, 1, 2, 3)
	slice4 = append(slice4, 4, 5)
	t.Log(slice3, slice4)

	x := copy(slice3, slice4)
	t.Log(slice3, x)
	//y:=copy(slice4,slice3)
	//t.Log(slice4,y)
}
