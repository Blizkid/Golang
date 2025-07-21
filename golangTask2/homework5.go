package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func run() {
	var counter int
	var mu sync.Mutex     // 可以保证同一时刻只有一个协程能修改
	var wg sync.WaitGroup //保证主协程等待所有子协程完成

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("最终计数器的值：", counter)

	var c Counter
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				c.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("最终计数器的值：", c.Get())

}

type Counter struct{ counter int64 }

func (c *Counter) Inc() { atomic.AddInt64(&c.counter, 1) }

func (c *Counter) Get() int64 { return atomic.LoadInt64(&c.counter) }
