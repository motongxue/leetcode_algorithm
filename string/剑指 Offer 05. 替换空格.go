package string

import "strings"

// TODO 注意字符串和[]byte间的相互转换，以及...使用方法
// Go中，三个点（...）表示将一个slice或数组打散为单个值
func replaceSpace(s string) string {
	if len(s) <= 1 {
		return s
	}
	ss := strings.Split(s, " ")
	result := []byte{}
	i := 0
	for ; i < len(ss)-1; i++ {
		result = append(result, []byte(ss[i]+"%20")...)
	}
	result = append(result, []byte(ss[i])...)
	return string(result)
}
