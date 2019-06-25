package leetcode

// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
// 例如，给定三角形：
//
// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
// 说明：
//
// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

// MinimumTotal120 三角形最小路径和
// dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
func MinimumTotal120(triangle [][]int) int {
	m := len(triangle)

	if m == 0 {
		return 0
	}

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		n := len(triangle[i])
		dp[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if i == 0 {
				dp[0][0] = triangle[0][0]

			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == n-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}

	minNum := dp[m-1][0]

	for i := 1; i < len(dp[m-1]); i++ {
		if minNum > dp[m-1][i] {
			minNum = dp[m-1][i]
		}
	}

	return minNum
}
