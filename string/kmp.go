package main

// KMP模板
func strStr(mubiao, moshi string) int {
	// i 表示目标串的下标，j表示模式串的下标
	j := -1
	next := make([]int, len(moshi))
	getNext(next, moshi)
	for i := 0; i < len(mubiao); i++ {
		for j >= 0 && mubiao[i] != moshi[j+1] {
			j = next[j]
		}
		if mubiao[i] == moshi[j+1] {
			j++
		}
		// 若j已经到达模式串的结尾，则返回
		if j == len(moshi)-1 {
			// 因为i才是在目标串中的下标，而j是模式串的下标，所以取i
			return i - len(moshi) + 1
		}
	}
	return -1
}
func getNext(next []int, moshi string) {
	j := -1
	next[0] = -1
	for i := 1; i < len(moshi); i++ {
		for j >= 0 && moshi[i] != moshi[j+1] {
			j = next[j]
		}
		if moshi[i] == moshi[j+1] {
			j++
		}
		next[i] = j
	}
}
