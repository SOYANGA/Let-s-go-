package share_mem

import (
	"sync"
	"testing"
	"time"
)

//共享内存counter多协程自增测试（不加锁） //4966不符合预期 非线程安全
func TestCounter1(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

//共享内存counter多协程自增测试（加锁） //5000符合预期 线程安全
func TestCounter2(t *testing.T) {
	//获取锁
	var mut sync.Mutex
	//共享资源
	counter := 0
	for i := 0; i < 5000; i++ {
		//协程并发执行
		go func() {
			//加锁
			mut.Lock()
			counter++
			//每个协程在执行完后进行解锁
			defer func() {
				mut.Unlock()
			}()
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

//共享内存counter多协程自增测试 + WaitGroup测试（主线程需要等待协程全部做完后再执行的方法）
func TestCounter3(t *testing.T) {
	//获取锁
	var mut sync.Mutex
	var wg sync.WaitGroup
	//共享资源
	counter := 0
	for i := 0; i < 5000; i++ {
		//协程并发执行每次启动一次执行，我们就需要添加一个我们需要等待的协程
		wg.Add(1)
		go func() {
			//加锁
			mut.Lock()
			counter++
			//每个协程在执行完后进行解锁并减少一次我们需要等待的协程
			defer func() {
				wg.Done()
				mut.Unlock()
			}()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
