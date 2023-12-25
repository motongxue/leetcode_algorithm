// 启动 2个groutine 2秒后取消， 第一个协程1秒执行完，第二个协程3秒执行完。
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func slp(ctx context.Context, timeout time.Duration) {
	defer wg.Done()
	select {
	case <-time.After(timeout):
		fmt.Printf("Sleeping for %+v\n", timeout)
	case <-ctx.Done():
		fmt.Printf("ctx cancel\n")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg.Add(2)
	go slp(ctx, time.Second)
	go slp(ctx, time.Second*3)
	wg.Wait()
	fmt.Println("Exiting main")
}
