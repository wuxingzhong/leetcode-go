/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high := 0, n1
	mid1, mid2 := 0, 0
	n := (n1 + n2 + 1) >> 1
	for low <= high {
		mid1 = low + (high-low)>>1
		mid2 = n - mid1
		if mid1 > 0 && nums1[mid1-1] > nums2[mid2] {
			high = mid1 - 1
		} else if mid1 != n1 && nums1[mid1] < nums2[mid2-1] {
			low = mid1 + 1
		} else {
			break
		}
	}

	leftMid, rightMid := 0, 0
	if mid1 == 0 {
		leftMid = nums2[mid2-1]
	} else if mid2 == 0 {
		leftMid = nums1[mid1-1]
	} else {
		leftMid = max(nums1[mid1-1], nums2[mid2-1])
	}

	if (n1+n2)%2 == 1 {
		return float64(leftMid)
	}

	if mid1 == n1 {
		rightMid = nums2[mid2]
	} else if mid2 == n2 {
		rightMid = nums1[mid1]
	} else {
		rightMid = min(nums1[mid1], nums2[mid2])
	}

	return float64(leftMid+rightMid) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

