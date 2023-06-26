package main

import (
	"fmt"
	"sync"
)

// 要求：交替打印A1B2C3D4E5…
func main() {
	charChan := make(chan struct{}, 1)
	digChan := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 26; i++ {
			<-charChan
			fmt.Printf("%c\n", 'A'+i)
			digChan <- struct{}{}
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 26; i++ {
			<-digChan
			fmt.Println(i)
			charChan <- struct{}{}
		}
		wg.Done()
	}()
	charChan <- struct{}{}
	wg.Wait()
}
