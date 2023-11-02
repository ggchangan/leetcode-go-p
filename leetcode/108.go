package leetcode

var SortedArrayToBST = sortedArrayToBST

func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(l, r int) *TreeNode
	helper = func(l, r int) *TreeNode {
		if r-l == 0 {
			return nil
		}
		if r-l == 1 {
			return &TreeNode{
				Val: nums[l],
			}
		}

		m := (l + r) / 2
		root := &TreeNode{
			Val: nums[m],
		}

		root.Left = helper(l, m)
		root.Right = helper(m+1, r)

		return root
	}

	return helper(0, len(nums))
}
