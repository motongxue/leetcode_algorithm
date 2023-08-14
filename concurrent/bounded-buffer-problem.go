package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 要求：开两个协程，一个生产者，生产10内的随机数，一个消费者，消费10个数结束
func main() {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			// 生成10以内的随机数
			ch <- rand.Intn(10)
		}
		close(ch)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("==========finished==========")
}
