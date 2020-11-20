/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return
}

// @lc code=end

