package csp

import (
	"fmt"
	"testing"
	"time"
)

//主要任务
func service() string {
	time.Sleep(time.Millisecond * 50)
	return "MainTask is Done"
}

//其他任务
func otherTask() {
	fmt.Println("Working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("OtherTask is done")
}

//串行执行两个任务 --- PASS: TestService (0.15s)
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

//对以上的service进行了简单的channel包装 异步化
func AsyncService() chan string {
	//阻塞式一放一取式，缺一就会阻塞（消息+另一方）
	retCh := make(chan string)

	//buffer channel且设定的容量是1，发送方放完消息就不会干等接收方去取（阻塞），而是会向下继续执行代码
	//retCh := make(chan string,1 )
	//启用另外的协程去运行，而不是阻塞当前协程去处理
	go func() {
		//取得service结果
		ret := service()
		fmt.Println("returned result.")
		//将结果放入channel中 只有接收端从channel中去取，此协程才会向下执行，反之则会被阻塞在这里
		retCh <- ret
		fmt.Println("Service exited")
	}()
	//返回channel
	return retCh
}

//引入channel异步化测试
func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}

//1.阻塞式channel
//Working on something else
//returned result.
//OtherTask is done
//MainTask is Done   channel中消息被取出，则才执行接下来的代码
//Service exited

//2. buffer channel 更高效的写法
//Working on something else
//returned result.
//Service exited   发送方无需阻塞的等待接收方去消费，直接执行接下来的代码
//OtherTask is done
//MainTask is Done
