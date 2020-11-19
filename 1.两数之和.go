/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, k}
		}
		m[v] = k
	}
	return nil
}

// @lc code=end

