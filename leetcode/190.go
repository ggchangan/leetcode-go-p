package leetcode

var ReverseBits = reverseBits

// 基础的位运算，使用与1加向右移位获取相应的位，然后将此位通过或运算加到相应的位上
// O(log n)
// O(1)
func reverseBits(num uint32) (ans uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		ans |= (num & 1) << (31 - i)
		num >>= 1
	}
	return
}
