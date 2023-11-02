package leetcode

var LongestConsecutive = longestConsecutive

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}

	maxL := 0
	for _, k := range nums {
		if m[k-1] { //skip
			continue
		}
		curL := 1
		for n := k + 1; m[n]; n++ {
			curL++
		}
		if curL > maxL {
			maxL = curL
		}
	}

	return maxL
}
