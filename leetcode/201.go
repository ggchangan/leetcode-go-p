package leetcode

var RangeBitwiseAnd = rangeBitwiseAnd

// 方法1
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left, right = left>>1, right>>1
		shift++
	}

	return left << shift
}

//方法2 n&(n-1)去掉最右面的1,直到left>=right, right就是公共前缀
