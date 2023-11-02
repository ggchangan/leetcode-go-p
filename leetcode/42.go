package leetcode

var Trap = trap

// 思路：求水的总面积，最开始时自己的思路是把水的分布看成一个个V，V中可以存水，企图通过查找V，并计算V中存水量计算。此方法是错误，因为当前列水能够存多少是由这个列左边最大值以及右边最大值决定的
// 可以将此题的结构，看成一个矩阵，可以安行求水，也可以按列求水
// 按行求水，类似于模拟算法，对于第i行，如果高度小于i，并且左右两边有大于height[i]的柱子可以积水，
// 按列求水，每一列水的高度，由此列左边最大值和右边最大值中的最小值决定，暴力解法对于每一列，都向前扫描求左最大值和想向后扫描求最大值
// 优化，左最大值left[i]和left[i-1]之间存在递推关系，可以提前计算；右最大值也是
func trap(height []int) (ans int) {
	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}
	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	l := len(height)
	if l <= 1 {
		return 0
	}
	leftMax := make([]int, l)
	for i := 1; i < l; i++ {
		leftMax[i] = maxHelper(leftMax[i-1], height[i-1])
	}
	rightMax := make([]int, l)
	for i := l - 2; i >= 0; i-- {
		rightMax[i] = maxHelper(rightMax[i+1], height[i+1])
	}

	for i := 0; i < l; i++ {
		minMax := minHelper(leftMax[i], rightMax[i])
		if height[i] < minMax {
			ans += minMax - height[i]
		}
	}

	return
}
