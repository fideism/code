package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup 一个异步结构体
type WaitGroup struct {
	workChan chan int
	wg       sync.WaitGroup
}

// NewPool 生成一个工作池, coreNum 限制
func NewPool(coreNum int) *WaitGroup {
	ch := make(chan int, coreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

// Add 添加
func (ap *WaitGroup) Add(num int) {
	fmt.Println("add", num)
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

func main() {
	work := NewPool(10)

	for i := 0; i < 50; i++ {
		work.Add(1)
		go testFunc(i, work)
	}

	fmt.Println("waiting...")
	work.Wait()

	fmt.Println("done")

	song()
}

func testFunc(i int, wg *WaitGroup) {
	defer wg.Done()
	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"), "output: ", i)
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"), "output: ", i, "done")
}

type Song struct {
	ID int
}

type SongWork struct {
	song chan Song
	wait sync.WaitGroup
}

func Pool(num int) *SongWork {
	ch := make(chan Song, num)

	return &SongWork{
		song: ch,
		wait: sync.WaitGroup{},
	}
}

func (s *SongWork) Add(num int) {
	for i := 0; i < num; i++ {
		s.song <- Song{ID: i}
		s.wait.Add(1)
	}
}

func (s *SongWork) Done() {
LOOP:
	for {
		select {
		case <-s.song:
			break LOOP

		}
	}

	s.wait.Done()
}

func (s *SongWork) Wait() {
	s.wait.Wait()
}

func song() {
	work := Pool(10)

	for i := 0; i < 50; i++ {
		work.Add(1)

		go testSong(Song{ID: i}, work)
	}

	work.Wait()
}

func testSong(song Song, s *SongWork) {
	defer s.Done()

	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"), "song: ", song.ID)
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"), "song: ", song.ID, "done")
}
