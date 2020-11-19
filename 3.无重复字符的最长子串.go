/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	rt, ans := -1, 0
	n := len(s)

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rt+1 < n && m[s[rt+1]] == 0 {
			m[s[rt+1]]++
			rt++
			ans = max(ans, rt-i+1)
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

