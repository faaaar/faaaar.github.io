package leetcode

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。
//
// 示例 1：
//
// 输入: "babad"
// 输出: "bab"
// 注意: "aba"也是一个有效答案。
// 示例 2：
//
// 输入: "cbbd"
// 输出: "bb"

// LongestPalindrome 最长回文子串
func LongestPalindrome(s string) string {
	l, r, slen := 0, 0, len(s)

	if slen < 2 {
		return s
	}

	for i := 0; i < slen; i++ {
		m, n := check(s, i, i, slen), check(s, i, i+1, slen)

		if m < n {
			m = n
		}

		if r-l+1 < m {
			// 此处根据当前子串长度和当前字符位置求出子串的左右边界的下标
			l = i - (m-1)/2
			r = i + m/2
		}
	}

	return s[l : r+1]
}

func check(s string, l, r, slen int) int {
	for l >= 0 && r < slen && s[l] == s[r] {
		l--
		r++
	}

	// 由于最后依次多处理了依次，所以需要还原l和r的值
	// 数组中下标a到下标b的长度为a-b+1 (a>b)
	// (r-1)-(l+1)+1 = r-l-1
	return r - l - 1
}
