# 第一个Go程序——Hello World

![1](https://github.com/SOYANGA/Let-s-go-/blob/master/picture/1.png)

## Show me the code

```go
package main  //包，表明代码所在模块(包)

import "fmt" //引入代码依赖

//功能实现
func main() {
	fmt.Println("Hello World")
}
```

## 应用程序的入口

1. **必须是main包：package main**
2. **必须是main方法：func main()**
3. **文件名不一定是main.go**

## 退出返回值

与其他主要编程语言的差异

### Go中main函数不支持任何返回值

### 通过os.Exit来返回状态

- ```go
  os.Exit(-1)
  ```



## 获取命令行参数

与其他主要编程语言的差异

### main函数不支持传入参数

### 在程序中直接通过os.Args获取命令行参数（返回的是数组形式的命令行参数）





> The master has failed more times than the beginner has tried.
>
> 大师失败的次数比初学者尝试的次数要多很多

## 编写测试程序

### 源码文件以_test结尾：xxx_test.go

### 测试方法名以Test开头：func TestXXX(t *testing.T) {…}

```go
package try_test

import "testing"

func TestFirstTry( t * testing.T) {
	t.Log("My first try!")
}
```



## 变量赋值

与其他主要编程语言的差异

### 赋值可以进行自动类型判断

### 在一个赋值语句中可以对多个变量进行同时赋值

```go
func TestFibList(t *testing.T) {
	//1.
	//var a int = 1
	//var b int = 1

	//2.
	//var(
	//	a int= 1
	//	b = 1
	//)

	//3.推荐 使用类型推断 直接给变量赋值
	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		//temp := a
		//a = b
		//b = temp + a
		a, b = b, a+b //对多个变量进行同时赋值
	}
}
```



交换两个变量的值

```go
/**
交换两个变量的值
*/
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//temp := a
	//a = b
	//b = temp
	a, b = b, a
	t.Log(a, b)
}
```



## 常量定义（初始化方式）

与其他主要编程语言的差异

### **快速设置连续值(常量)**

```go
package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)
/**
分别用二进制的三个位标识，三种状态
 */
const (
	Readable   = 1 << iota //可读
	Writable               //可写
	Executable             //可执行
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	b := 1 //0001
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
}

```

