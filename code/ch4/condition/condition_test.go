package condition

import "testing"

/*
go语言的for循环使用
*/
func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	/*
	 可以提供多返回值时对方法错误的判断以及处理
	*/
	//if v, err := someFun(); err == nil {
	//	t.Log("1 == 1")
	//} else {
	//	t.Log("1 == 1")
	//}
}

/*
go语言的switch case使用
*/
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")

		}
	}
}

//condition_test.go:30: Even
//condition_test.go:32: odd
//condition_test.go:30: Even
//condition_test.go:32: odd
//condition_test.go:34: it is not 0-3

/*
利用switch case 简化复杂在的if else操作
*/
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i&1 == 0:
			t.Log("Even")
		case i&1 == 1:
			t.Log("odd")
		default:
			t.Log("unknown")
		}
	}
}

//condition_test.go:50: Even
//condition_test.go:52: odd
//condition_test.go:50: Even
//condition_test.go:52: odd
//condition_test.go:50: Even
