package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

//func (e *Employee) String() string {
//	fmt.Printf("Name's Address is %x", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

func (e Employee) String() string {
	fmt.Printf("Name's Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

/*
测试声明对方法定义的影响
1.func (e Employee) String() string {} 发生了拷贝不推荐使用 有空间复制的开销
Name's Address is c000056550
Name's Address is c000056580
实例的成员是经过复制后的成员，即成员指向不同的空间

2.func (e *Employee) String() string {} （推荐使用避免拷贝）
Name's Address is c000056550
Name's Address is c000056550
则实例的成员并没有被对象所拷贝，两个方法中的实例的成员是指向相同的空间
*/
func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20} //通过实例的指针访问成员不需要使用-> 使用 .
	fmt.Printf("Name's Address is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.String()) //ID:0-Name:Bob-Age:20
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e)             //ID:0-Name:Bob-Age:20
	t.Logf("e is %T", e) //e is encapsulation.Employee

	e1 := Employee{Name: "Mike", Age: 30}
	t.Log(e1.Id) //
	t.Log(e1)    //ID:-Name:Mike-Age:30

	e2 := new(Employee) //返回的是指针 相当于*Employee
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e2)              //ID:2-Name:Rose-Age:22
	t.Logf("e2 is %T", e2) //e2 is *encapsulation.Employee
	t.Logf("&e is %T", &e) //&e is *encapsulation.Employee

}
