package pool

import (
	"sync"
)

// WaitGroup 一个异步结构体
type WaitGroup struct {
	workChan chan int
	wg       sync.WaitGroup
}

// NewPool 生成一个工作池, coreNum 限制
func NewPool(num int) *WaitGroup {
	ch := make(chan int, num)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

// Add 添加
func (ap *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		ap.workChan <- i
		ap.wg.Add(1)
	}
}

// Done 完结
func (ap *WaitGroup) Done() {
LOOP:
	for {
		select {
		case <-ap.workChan:
			break LOOP
		}
	}
	ap.wg.Done()
}

// Wait 等待
func (ap *WaitGroup) Wait() {
	ap.wg.Wait()
}
