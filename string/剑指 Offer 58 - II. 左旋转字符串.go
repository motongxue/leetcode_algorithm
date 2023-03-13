func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse(b, 0, n-1)
	reverse(b, n, len(s)-1)
	reverse(b, 0, len(s)-1)
	return string(b)
}
func reverse(b []byte, l, r int) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}