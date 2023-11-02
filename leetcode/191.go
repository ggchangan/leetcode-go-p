package leetcode

var HammingWeight = hammingWeight

// 统计数字中每一位是否为1
// 将1向上移位，与数字相与，如果结果大于0，则相应位为1，否则为0
func hammingWeight(num uint32) (ans int) {
	for i := 0; i < 32; i++ {
		if num&(1<<i) > 0 {
			ans++
		}
	}

	return
}

//TOOD n & (n - 1) 会将最低位的1变成0，每次运算变1位，可计算此运算将n变成0，运算的次数
