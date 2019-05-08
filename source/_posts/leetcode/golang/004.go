package leetcode

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。
//
// 示例 1:
//
// nums1 = [1, 3]
// nums2 = [2]
//
// 则中位数是 2.0
// 示例 2:
//
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// 则中位数是 (2 + 3)/2 = 2.5

// FindMedianSortedArrays004 寻找两个有序数组的中位数
func FindMedianSortedArrays004(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 > n2 {
		return FindMedianSortedArrays004(nums2, nums1)
	}

	k := (n1 + n2 + 1) / 2
	l := 0
	r := n1

	for l < r {
		m1 := l + (r-l)/2
		m2 := k - m1

		if nums1[m1] < nums2[m2-1] {
			l = m1 + 1
		} else {
			r = m1
		}
	}

	m1 := l
	m2 := k - m1

	var t1, t2, c1, c2 int

	if m1 <= 0 {
		t1 = IntMin
	} else {
		t1 = nums1[m1-1]
	}

	if m2 <= 0 {
		t2 = IntMin
	} else {
		t2 = nums2[m2-1]
	}

	if t1 >= t2 {
		c1 = t1
	} else {
		c1 = t2
	}

	if (n1+n2)%2 == 1 {
		return float64(c1)
	}

	if m1 >= n1 {
		t1 = IntMax
	} else {
		t1 = nums1[m1]
	}

	if m2 >= n2 {
		t2 = IntMax
	} else {
		t2 = nums2[m2]
	}

	if t1 <= t2 {
		c2 = t1
	} else {
		c2 = t2
	}

	return (float64(c1 + c2)) / 2
}
