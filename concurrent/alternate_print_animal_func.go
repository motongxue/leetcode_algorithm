package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 不采用匿名函数形式
// 定义 3 个函数分别打印 cat、dog、fish，要求每个函数都要起一个 goroutine，按照 cat、dog、fish 顺序打印在屏幕上 100 次。
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	go printCat(catChan, dogChan, &wg)
	go printDog(dogChan, fishChan, &wg)
	go printFish(fishChan, catChan, &wg)
	catChan <- struct{}{}
	wg.Wait()
}
func printCat(catChan, dogChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-catChan
		fmt.Println("cat" + strconv.Itoa(i))
		dogChan <- struct{}{}
	}
}
func printDog(dogChan, fishChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-dogChan
		fmt.Println("dog" + strconv.Itoa(i))
		fishChan <- struct{}{}
	}
}
func printFish(fishChan, catChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-fishChan
		fmt.Println("fish" + strconv.Itoa(i))
		catChan <- struct{}{}
	}
}
