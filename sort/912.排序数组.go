/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
    // 将待排序数组构造成一个大根堆
    heapInsert(nums)
    size := len(nums)
    for size > 1 {
        // 固定一个最大值，将剩余的数再构造成一个大根堆
        swap(nums, 0, size - 1)
        size--
        heaptify(nums, 0, size)
    }
    return nums
}

// 构建大根堆
func heapInsert(nums []int) {
    for i := 0; i < len(nums); i++ {
        curIdx := i
        fatherIdx := (curIdx - 1) / 2
        
        for nums[curIdx] > nums[fatherIdx] {
            swap(nums, curIdx, fatherIdx)
            curIdx = fatherIdx
            fatherIdx = (curIdx - 1) / 2
        }
    }
}

func heaptify(nums []int, index, size int) {
    left := index * 2 + 1
    right := index * 2 + 2
    for left < size {
        // 从左中右中找到最大的值
        maxIdx := left
        if right < size && nums[right] > nums[left] {
            maxIdx = right
        }
        if nums[index] > nums[maxIdx] {
            maxIdx = index
        }
        // 如果 nums[index] 已经最大的，说明以index为根的子树本身已经稳定
        if maxIdx == index {
            break
        }
        swap(nums, maxIdx, index)
        // 开始从上往下，将剩余的数调整为大根堆
        index = maxIdx
        left = index * 2 + 1
        right = index * 2 + 2
    }
}

func swap(nums []int, i, j int) {
    nums[i], nums[j] = nums[j], nums[i]
}
// @lc code=end

