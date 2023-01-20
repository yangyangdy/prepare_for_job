package main
import (
	"sync"
	"time"
)
// 当写锁阻塞时，新的读锁是无法申请的

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}