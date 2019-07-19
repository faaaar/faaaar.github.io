package leetcode

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
//
// 示例:
//
// 输入: 3
// 输出:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// 解释:
// 以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

// GenerateTrees095Recursive 不同的二叉搜索树 II 递归解法
func GenerateTrees095Recursive(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return genTree095(1, n)
}

func genTree095(start, end int) []*TreeNode {
	list := []*TreeNode{}
	if start > end {
		return append(list, nil)
	}

	if start == end {
		return append(list, &TreeNode{start, nil, nil})
	}

	for i := start; i <= end; i++ {
		leftList := genTree095(start, i-1)
		rightList := genTree095(i+1, end)

		for _, left := range leftList {
			for _, right := range rightList {
				list = append(list, &TreeNode{i, left, right})
			}
		}
	}

	return list
}

// GenerateTrees095DP 不同的二叉搜索树 II 动态规划解法
func GenerateTrees095DP(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{}

	if n == 0 {
		return dp[0]
	}

	dp[0] = append(dp[0], nil)

	// 从1到n，逐步推算n的树的数量
	for i := 1; i <= n; i++ {
		dp[i] = []*TreeNode{}
		// 在i的前提下，根节点为j时的树的情况
		for j := 1; j <= i; j++ {
			// dp[j-1]左子树数组
			for _, left := range dp[j-1] {
				// dp[i-j]右子树数组
				for _, right := range dp[i-j] {
					root := &TreeNode{j, left, clone095(right, j)}

					dp[i] = append(dp[i], root)
				}
			}
		}
	}

	return dp[n]
}

func clone095(treeNode *TreeNode, offset int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	return &TreeNode{
		treeNode.Val + offset,
		clone095(treeNode.Left, offset),
		clone095(treeNode.Right, offset),
	}
}
