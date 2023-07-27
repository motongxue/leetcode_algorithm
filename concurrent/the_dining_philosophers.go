// 哲学家进餐问题，有n个人，保证不出现死锁
package main

import (
	"fmt"
	"sync"
)

var (
	ch = make(chan struct{}, 4)
	wg sync.WaitGroup
)

func main() {
	n := 10
	wg.Add(n)
	go philosopher1()
	go philosopher2()
	go philosopher3()
	go philosopher4()
	go philosopher5()
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	wg.Wait()
}
func philosopher1() {
	for {
		<-ch
		fmt.Println("philosopher1 is eating")
		wg.Done()
	}
}

func philosopher2() {
	for {
		<-ch
		fmt.Println("philosopher2 is eating")
		wg.Done()
	}
}

func philosopher3() {
	for {
		<-ch
		fmt.Println("philosopher3 is eating")
		wg.Done()
	}
}

func philosopher4() {
	for {
		<-ch
		fmt.Println("philosopher4 is eating")
		wg.Done()
	}
}

func philosopher5() {
	for {
		<-ch
		fmt.Println("philosopher5 is eating")
		wg.Done()
	}
}
