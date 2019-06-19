package leetcode

// 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
//
// 示例：
//
// 给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
//
// sumRange(0, 2) -> 1
// sumRange(2, 5) -> -1
// sumRange(0, 5) -> -3
// 说明:
//
// 你可以假设数组不可变。
// 会多次调用 sumRange 方法。

// NumArray303 数组
// Nums 源数据
// Sum 索引0-i的和
type NumArray303 struct {
	Nums []int
	Sum  []int
}

// Constructor303 数组构造函数
func Constructor303(nums []int) NumArray303 {
	l := len(nums)
	sum := make([]int, l)
	sum[0] = nums[0]

	for i := 1; i < l; i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	return NumArray303{
		Sum: sum,
	}
}

// SumRange303 区域和检索 - 数组不可变
func (numarray *NumArray303) SumRange303(i int, j int) int {
	if i == 0 {
		return numarray.Sum[j]
	}

	return numarray.Sum[j] - numarray.Sum[i-1]
}

/**
 * Your NumArray303 object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
