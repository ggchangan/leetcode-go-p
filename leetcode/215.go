package leetcode

import "math/rand"

var FindKthLargest = findKthLargest

// https://zh.mojotv.cn/algorithm/golang-quick-sort
func findKthLargest1(nums []int, k int) int {
	partition := func(l, r int) int {
		if r == l {
			return l
		}

		i := rand.Intn(r-l) + l //随机选取其中一个元素和最左边的替换
		nums[l], nums[i] = nums[i], nums[l]

		pivot := nums[l]
		for l < r {
			//注意=
			for l < r && nums[r] >= pivot {
				r--
			}

			nums[l] = nums[r]

			for l < r && nums[l] <= pivot {
				l++
			}

			nums[r] = nums[l]
		}
		nums[l] = pivot
		return l
	}

	numsL := len(nums)
	target := numsL - k
	l, r := 0, numsL-1
	for {
		pivot := partition(l, r)
		if target == pivot {
			return nums[target]
		} else if target > pivot {
			l = pivot + 1
		} else {
			r = pivot - 1
		}
	}
}

func heapify(nums []int, n, i int) {
	largest := i
	lson, rson := 2*i+1, 2*i+2
	if lson < n && nums[lson] > nums[i] {
		largest = lson
	}
	if rson < n && nums[rson] > nums[largest] {
		largest = rson
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, n, largest)
	}

	return
}

// for i , parent = (i-1)/2,最后一个节点的父节点为有孩子的第一个节点，其下标i=(n-1-1)/2 = n/2 -1
func buildMaxHeap(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

// var HeapSort func(nums []int)
var HeapSort = heapSort

func heapSort(nums []int) {
	buildMaxHeap(nums)

	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}
}

func findKthLargest(nums []int, k int) int {
	buildMaxHeap(nums)

	n := len(nums)
	for i := n - 1; i >= n-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	return nums[n-k]
}
