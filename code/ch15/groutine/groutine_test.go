package groutine__test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}

/*
go协程加入调用得到函数，并不是传参调用则就会产生共享内存问题，则就需要加锁
*/
