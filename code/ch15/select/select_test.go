package select__test

import (
	"fmt"
	"testing"
	"time"
)

//主要任务
func service() string {
	time.Sleep(time.Millisecond * 500)
	return "MainTask is Done"
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

//超时机制测试避免持续等待（阻塞）
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
