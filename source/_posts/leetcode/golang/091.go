package leetcode

// 一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
// 示例 1:
//
// 输入: "12"
// 输出: 2
// 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
// 示例 2:
//
// 输入: "226"
// 输出: 3
// 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

// NumDecodings091 解码方法
// 注意对0的处理
func NumDecodings091(s string) int {
	if s[0] == '0' {
		return 0
	}

	l := len(s)

	if l == 1 {
		return 1
	}

	dp := make([]int, l+1)
	dp[0] = 1

	for i := 1; i <= l; i++ {
		if s[i-1] == '0' {
			dp[i] = 0
			if s[i-2] < 1 || s[i-2] > 2 {
				return 0
			}

		} else {
			dp[i] = dp[i-1]
		}

		if i-1 > 0 && ((s[i-2] == '1') || (s[i-2] == '2' && s[i-1] <= '6')) {
			dp[i] += dp[i-2]
		}
	}

	return dp[l]
}
