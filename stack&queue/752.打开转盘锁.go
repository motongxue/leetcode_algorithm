/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start
// ===========================================================
// 利用BFS进行求解
type Set map[string]struct{}
// 如果已经存在，则添加失败
func (s Set) Add(in string) bool {
	_, exists := s[in]
	if !exists {
		s[in] = struct{}{}
	}
	return !exists
}
func openLock(deadends []string, target string) int {
    // 如果终点已经是起点
    if target == "0000" {
		return 0
	}

    step := 0
    st := Set{}
    for _, v := range deadends {
        st.Add(v)
    }

    // 防止起点为死锁点
    if _, exist := st["0000"]; exist {
        return -1
    }
    q := []string{"0000"}
    
    for len(q) != 0 {
        step++
        size := len(q)

        for i := 0; i < size; i++ {
            cur := q[0]
            q = q[1:]

            for j := 0; j < 4; j++ {
                if flg := process(cur, target, j, 1, &st, &q); flg {
                    return step
                }
                if flg := process(cur, target, j, -1, &st, &q); flg {
                    return step
                }
            }
        }
    }
    return -1
}
// 如果已经找到了，则返回true
func process(in, target string, idx, method int, st *Set, q *[]string) bool {
    tmp := changeString(in, idx, method)
    if tmp == target {
        return true
    }
    
    if st.Add(tmp) {
        *q = append(*q, tmp)
    }
    return false
}
func changeString(in string, idx, method int) string {
    ch := []byte(in)

    if method == 1 {
        if ch[idx] == '9' {
            ch[idx] = '0'
        }else {
            ch[idx] += 1
        }
    }else {
        if ch[idx] == '0' {
            ch[idx] = '9'
        }else {
            ch[idx] -= 1
        }
    }
    return string(ch)
}

// @lc code=end

