package leetcode

var SingleNumber2 = singleNumber2

// 取所有数字对应的位，统计1的个数，并对3取余，得到对应位的数字，并与到对应位中
// 注意，题目中数字的范围是32位有符号整数
// go语言中，int在64位机器中，数字是有64位的，这里需要将参与运算的数字转化为int32
func singleNumber2(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		bit := int32(0)
		for j := 0; j < len(nums); j++ {
			bit += (int32(nums[j]) >> i) & 1
		}
		ans |= (bit % 3) << i
	}

	return int(ans)
}
