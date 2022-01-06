package main

import (
	"fmt"
	"time"

	"github.com/fideism/code/pool"
)

func main() {
	work := pool.NewPool(10)

	for i := 0; i < 10; i++ {
		work.Add(1)
		go testPoolFunc(i, work)
	}

	fmt.Println("waiting...")
	work.Wait()
	fmt.Println("done")
}

func testPoolFunc(i int, wg *pool.WaitGroup) {
	defer wg.Done()
	fmt.Println(time.Now(), "output: ", i)
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now(), "output: ", i, "done")
}
