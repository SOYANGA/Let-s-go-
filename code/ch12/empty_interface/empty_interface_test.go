package empty_interface

import (
	"fmt"
	"testing"
)

/*
接口查询 和类型查询
 */
func DoSomething(p interface{}) {
	//接口查询
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("string", s)
	//	return
	//}
	//fmt.Println("Unknown Type")

	//类型查询
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown Type")
	}
}

/*
接口查询
*/
func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)    //Integer 10
	DoSomething("100") //string 100
}
