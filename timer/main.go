package main

import (
	"fmt"
	"time"
)

func main() {
	// 设置时间间隔
	t1 := time.NewTimer(time.Second * 3)
	for {
		select {
		case <-t1.C:
			timerDeal()
			t1.Reset(time.Second * 3)
		}
	}
}

func timerDeal() {
	fmt.Println(time.Now())
}
