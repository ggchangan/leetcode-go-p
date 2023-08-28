package main

import "fmt"

//88
//func main() {
//	fmt.Println("Hello world!")
//	nums1 := []int{1, 2, 3, 0, 0, 0}
//	nums2 := []int{2, 5, 6}
//	merge(nums1, 3, nums2, 3)
//	fmt.Println(nums1)
//
//	nums1 = []int{1}
//	nums2 = []int{}
//	merge(nums1, 1, nums2, 0)
//	fmt.Println(nums1)
//
//	nums1 = []int{0}
//	nums2 = []int{1}
//	merge(nums1, 0, nums2, 1)
//	fmt.Println(nums1)
//}

// 27
//func main() {
//	nums := []int{3, 2, 2, 3}
//	val := 3
//	i := removeElement(nums, val)
//	fmt.Println(i)
//	fmt.Println(nums)
//
//	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
//	val = 2
//	i = removeElement(nums, val)
//	fmt.Println(i)
//	fmt.Println(nums)
//	return
//}

// 26

//func main() {
//	nums := []int{1, 1, 2}
//	l := removeDuplicates(nums)
//	fmt.Println(nums)
//	fmt.Println(l)
//
//	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
//	l = removeDuplicates(nums)
//	fmt.Println(nums)
//	fmt.Println(l)
//}

// 80
//func main() {
//	nums := []int{1, 1, 1, 2, 2, 3}
//	l := removeDuplicates80(nums)
//	fmt.Println(nums)
//	fmt.Println(l)
//
//	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
//	l = removeDuplicates80(nums)
//	fmt.Println(nums)
//	fmt.Println(l)
//}

// 169
//func main() {
//	nums := []int{3, 2, 3}
//	number := majorityElement(nums)
//	fmt.Println(number)
//
//	nums = []int{2, 2, 1, 1, 1, 2, 2}
//	number = majorityElement(nums)
//	fmt.Println(number)
//}

// 189
//func main() {
//	nums := []int{1, 2, 3, 4, 5, 6, 7}
//	k := 3
//	rotate(nums, k)
//	fmt.Println(nums)
//
//	nums = []int{-1, -100, 3, 99}
//	k = 2
//
//	rotate(nums, k)
//	fmt.Println(nums)
//
//	nums = []int{-1}
//	k = 2
//
//	rotate(nums, k)
//	fmt.Println(nums)
//}

// 55
//func main() {
//	nums := []int{2, 3, 1, 1, 4}
//	fmt.Println(canJump(nums))
//
//	nums = []int{3, 2, 1, 0, 4}
//	fmt.Println(canJump(nums))
//}

// 70
//func main() {
//	fmt.Println(climbStairs(1))
//	fmt.Println(climbStairs(2))
//	fmt.Println(climbStairs(3))
//	fmt.Println(climbStairs(4))
//	fmt.Println(climbStairs(5))
//	fmt.Println(climbStairs(6))
//}

// 22
//func main() {
//	fmt.Println(generateParenthesis(1))
//	fmt.Println(generateParenthesis(2))
//	fmt.Println(generateParenthesis(3))
//}

// 17
func main() {
	digits := "2,3"
	fmt.Println(letterCombinations(digits))

	digits = ""
	fmt.Println(letterCombinations(digits))

	digits = "2"
	fmt.Println(letterCombinations(digits))

}
