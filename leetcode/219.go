package leetcode

var ContainsNearbyDuplicate = containsNearbyDuplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]bool{}
	p, q := 0, 0
	for ; p < len(nums) && p < k+1; p++ {
		if _, ok := m[nums[p]]; ok {
			return true
		}
		m[nums[p]] = true
	}

	for ; p < len(nums); p++ {
		delete(m, nums[q])
		if _, ok := m[nums[p]]; ok {
			return true
		}
		m[nums[p]] = true

		q++
	}

	return false
}
