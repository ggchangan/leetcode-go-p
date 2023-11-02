package leetcode

var Candy = candy

// 感悟：开始时，完全没有思路，想着左右比较，
// 数据结构加算法，整个是线性结构，孩子左右两边是线，构成双链表，对双链表进行遍历
// 做题时，可以尝试抽象出来数据、对象的结构，然后在考虑基于这样结构的算法

// 题目要求相邻两个孩子评分更高的孩子获取更多的糖果
// 思路：numbers存储每个孩子最少糖果数，
// 每个孩子有左右两个相邻的孩子，从左向右看，if ratings[i] > ratings[i-1] {numbers1[i]=numbers1[i-1]+1} else {numbers1[i]=1},
// 从右往左看， if ratings[i] > ratings[i+1] {numbers2[i-1]=numbers2[i-1]+1} else {numbers2[i]=1}
// 同时满足两个条件，则numbers[i] = max(numbers1, numbers2)
// 具体实现时，numbers2[i]只与上一个numbers2[i-1]有关，同时使用使用numbers1代替numbers
func candy(ratings []int) (ans int) {
	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}
	l := len(ratings)
	numbers := make([]int, l)
	for i := range ratings {
		if i > 0 && ratings[i] > ratings[i-1] {
			numbers[i] = numbers[i-1] + 1
		} else {
			numbers[i] = 1
		}
	}

	var right int
	for i := l - 1; i >= 0; i-- {
		if i < l-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		numbers[i] = maxHelper(numbers[i], right)
		ans += numbers[i]
	}

	return
}
