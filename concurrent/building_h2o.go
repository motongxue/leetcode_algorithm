// 1117. H2O 生成
// 现在有两种线程，氧 oxygen 和氢 hydrogen，你的目标是组织这两种线程来产生水分子。
// 存在一个屏障（barrier）使得每个线程必须等候直到一个完整水分子能够被产生出来。
// 氢和氧线程会被分别给予 releaseHydrogen 和 releaseOxygen 方法来允许它们突破屏障。
// 这些线程应该三三成组突破屏障并能立即组合产生一个水分子。
// 你必须保证产生一个水分子所需线程的结合必须发生在下一个水分子产生之前。

// 换句话说:
// 如果一个氧线程到达屏障时没有氢线程到达，它必须等候直到两个氢线程到达。
// 如果一个氢线程到达屏障时没有其它线程到达，它必须等候直到一个氧线程和另一个氢线程到达。
// 书写满足这些限制条件的氢、氧线程同步代码。

// 示例 1:
// 输入: water = "HOH"
// 输出: "HHO"
// 解释: "HOH" 和 "OHH" 依然都是有效解。

// 示例 2:
// 输入: water = "OOHHHH"
// 输出: "HHOHHO"
// 解释: "HOHHHO", "OHHHHO", "HHOHOH", "HOHHOH", "OHHHOH", "HHOOHH", "HOHOHH" 和 "OHHOHH" 依然都是有效解。
package main

import (
	"fmt"
	"sync"
)

func main() {
	water := "OOHHHHOOHHHHHHOHOH"

	OChan := make(chan struct{}, 1)
	HChan := make(chan struct{}, 2)
	defer close(OChan)
	defer close(HChan)

	wg := sync.WaitGroup{}
	wg.Add(len(water) / 3)

	go func() {
		for {
			// 一个O原子需要两个H原子
			<-OChan
			<-HChan
			<-HChan
			fmt.Print("HHO")
			wg.Done()
		}
	}()

	for i := 0; i < len(water); i++ {
		if water[i] == 'O' {
			OChan <- struct{}{}
		} else {
			HChan <- struct{}{}
		}
	}

	wg.Wait()
}
