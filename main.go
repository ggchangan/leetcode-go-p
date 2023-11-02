package main

import (
	"fmt"
	"leetcode-go-p/leetcode/backtrace"
)

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
//func main() {
//	digits := "23"
//	fmt.Println(leetcode.LetterCombinations(digits))
//
//	digits = ""
//	fmt.Println(leetcode.LetterCombinations(digits))
//
//	digits = "2"
//	fmt.Println(leetcode.LetterCombinations(digits))
//
//}

// 51
//func main() {
//	fmt.Println(solveNQueens(4))
//	fmt.Println(solveNQueens(1))
//}

// 433
//
//	func main() {
//		startGene := "AACCGGTT"
//		endGene := "AACCGGTA"
//		bank := []string{"AACCGGTA"}
//		fmt.Println(minMutation(startGene, endGene, bank))
//
//		startGene = "AACCGGTT"
//		endGene = "AAACGGTA"
//		bank = []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
//		fmt.Println(minMutation(startGene, endGene, bank))
//
//		startGene = "AAAAACCC"
//		endGene = "AACCCCCC"
//		bank = []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
//		fmt.Println(minMutation(startGene, endGene, bank))
//
//		startGene = "AAAAAAAT"
//		endGene = "CCCCCCCC"
//		bank = []string{"AAAAAAAC", "AAAAAAAA", "CCCCCCCC"}
//		fmt.Println(minMutation(startGene, endGene, bank))
//	}

//func main() {
//	beginWord := "hit"
//	endWord := "cog"
//	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
//
//	fmt.Println(ladderLength(beginWord, endWord, wordList))
//
//	beginWord = "hit"
//	endWord = "cog"
//	wordList = []string{"hot", "dot", "dog", "lot", "log"}
//	fmt.Println(ladderLength(beginWord, endWord, wordList))
//
//	//开始这个问题不过，并不是因为leetcode多线程问题，而是其中一个小函数逻辑有问题
//	beginWord = "hog"
//	endWord = "cog"
//	wordList = []string{"cog"}
//
//	fmt.Println(ladderLength(beginWord, endWord, wordList))
//}

//func main() {
//	grid := [][]byte{
//		[]byte{'1', '1', '1', '1', '0'},
//		[]byte{'1', '1', '0', '1', '0'},
//		[]byte{'1', '1', '0', '0', '0'},
//		[]byte{'0', '0', '0', '0', '0'},
//	}
//	fmt.Println(numIslands(grid))
//
//	fmt.Println("----------------------------------------------------------------")
//	grid = [][]byte{
//		[]byte{'1', '1', '0', '0', '0'},
//		[]byte{'1', '1', '0', '0', '0'},
//		[]byte{'0', '0', '1', '0', '0'},
//		[]byte{'0', '0', '0', '1', '1'},
//	}
//	fmt.Println(numIslands(grid))
//}

//func main() {
//	fmt.Println(uniquePaths(3, 7))
//	fmt.Println(uniquePaths(3, 2))
//	fmt.Println(uniquePaths(7, 3))
//	fmt.Println(uniquePaths(3, 3))
//}

//func main() {
//	obstacleGrid := [][]int{
//		[]int{0, 0, 0},
//		[]int{0, 1, 0},
//		[]int{0, 0, 0},
//	}
//	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
//	obstacleGrid = [][]int{
//		[]int{0, 1},
//		[]int{0, 0},
//	}
//	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
//}

//func main() {
//	fmt.Println(longestCommonSubsequence("abcde", "ace"))
//	fmt.Println(longestCommonSubsequence("abc", "abc"))
//	fmt.Println(longestCommonSubsequence("abc", "def"))
//}

//func main() {
//	triangle := [][]int{
//		{2},
//		{3, 4},
//		{6, 5, 7},
//		{4, 1, 8, 3},
//	}
//	fmt.Println(minimumTotal(triangle))
//
//	triangle = [][]int{
//		{-10},
//	}
//	fmt.Println(minimumTotal(triangle))
//}

//func main() {
//	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
//	fmt.Println(maxSubArray(nums))
//
//	nums = []int{1}
//	fmt.Println(maxSubArray(nums))
//
//	nums = []int{5, 4, -1, 7, 8}
//	fmt.Println(maxSubArray(nums))
//}

//func main() {
//	nums := []int{2, 3, -2, 4}
//	fmt.Println(maxProduct(nums))
//	fmt.Println(maxProduct2(nums))
//
//	nums = []int{-2, 0, -1}
//	fmt.Println(maxProduct(nums))
//	fmt.Println(maxProduct2(nums))
//}

//func main() {
//	coins := []int{1, 2, 5}
//	amount := 11
//	fmt.Println(coinChange(coins, amount))
//
//	coins = []int{2}
//	amount = 3
//	fmt.Println(coinChange(coins, amount))
//
//	coins = []int{1}
//	amount = 0
//	fmt.Println(coinChange(coins, amount))
//}

//func main() {
//	nums := []int{1, 2, 3, 1}
//	fmt.Println(rob(nums))
//
//	nums = []int{2, 7, 9, 3, 1}
//	fmt.Println(rob(nums))
//}

//func main() {
//	nums := []int{2, 3, 2}
//	fmt.Println(rob2(nums))
//
//	nums = []int{1, 2, 3, 1}
//	fmt.Println(rob2(nums))
//
//	nums = []int{1, 2, 3}
//	fmt.Println(rob2(nums))
//
//	nums = []int{1, 2}
//	fmt.Println(rob2(nums))
//}

//func main() {
//	s := "(()"
//	fmt.Println(longestValidParentheses(s))
//
//	s = ")()())"
//	fmt.Println(longestValidParentheses(s))
//
//	s = ""
//	fmt.Println(longestValidParentheses(s))
//}

//func main() {
//	grid := [][]int{
//		{1, 3, 1},
//		{1, 5, 1},
//		{4, 2, 1},
//	}
//	fmt.Println(minPathSum(grid))
//
//	grid = [][]int{
//		{1, 2, 3},
//		{4, 5, 6},
//	}
//	fmt.Println(minPathSum(grid))
//}

//func main() {
//	fmt.Println(minDistance("horse", "ros"))
//	fmt.Println(minDistance("intention", "execution"))
//	fmt.Println(minDistance("pneumonoultramicroscopicsilicovolcanoconiosis", "ultramicroscopically"))
//	fmt.Println(minDistance("sea", "eat"))
//}

//func main() {
//	fmt.Println(numDecodings("12"))
//	fmt.Println(numDecodings("226"))
//	fmt.Println(numDecodings("06"))
//	fmt.Println(numDecodings("2101"))
//	fmt.Println(numDecodings("2611055971756562"))
//}

//func main() {
//s := "leetcode"
//wordDict := []string{"leet", "code"}

//fmt.Println(wordBreak(s, wordDict))

//s = "applepenapple"
//wordDict = []string{"apple", "pen"}

//fmt.Println(wordBreak(s, wordDict))

//s = "catsandog"
//wordDict = []string{"cats", "dog", "sand", "and", "cat"}

//fmt.Println(wordBreak(s, wordDict))

//s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
//wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
//fmt.Println(wordBreak(s, wordDict))
//}

//func main() {
//	s := "A man, a plan, a canal: Panama"
//	fmt.Println(leetcode.IsPalindrome(s))
//
//	s = "race a car"
//	fmt.Println(leetcode.IsPalindrome(s))
//
//	s = "  "
//	fmt.Println(leetcode.IsPalindrome(s))
//
//	s = ""
//	fmt.Println(leetcode.IsPalindrome(s))
//}

//func main() {
//s := "abc"
//t := "ahbgdc"

//fmt.Println(leetcode.IsSubsequence(s, t))
//s = "axc"
//t = "ahbgdc"
//fmt.Println(leetcode.IsSubsequence(s, t))

//s = "aaaaaa"
//t = "bbaaaa"
//fmt.Println(leetcode.IsSubsequence(s, t))
//}

//func main() {
//	numbers := []int{2, 7, 11, 15}
//	target := 9
//	fmt.Println(leetcode.TwoSum(numbers, target))
//
//	numbers = []int{2, 3, 4}
//	target = 6
//	fmt.Println(leetcode.TwoSum(numbers, target))
//
//	numbers = []int{-1, 0}
//	target = -1
//	fmt.Println(leetcode.TwoSum(numbers, target))
//
//	numbers = []int{5, 25, 75}
//	target = 100
//	fmt.Println(leetcode.TwoSum(numbers, target))
//}

//func main() {
//	nums := []int{-1, 0, 1, 2, -1, -4}
//	fmt.Println(leetcode.ThreeSum(nums))
//
//	nums = []int{0, 1, 1}
//	fmt.Println(leetcode.ThreeSum(nums))
//
//	nums = []int{0, 0, 0}
//	fmt.Println(leetcode.ThreeSum(nums))
//}

//func main() {
//	nums := []int{2, 3, 1, 2, 4, 3}
//	target := 7
//	fmt.Println(leetcode.MinSubArrayLen(target, nums))
//
//	nums = []int{1, 4, 4}
//	target = 4
//	fmt.Println(leetcode.MinSubArrayLen(target, nums))
//
//	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
//	target = 11
//	fmt.Println(leetcode.MinSubArrayLen(target, nums))
//}

//func main() {
//	s := "abcabcbb"
//	fmt.Println(leetcode.LengthOfLongestSubstring(s))
//
//	s = "bbbbb"
//	fmt.Println(leetcode.LengthOfLongestSubstring(s))
//
//	s = "pwwkew"
//	fmt.Println(leetcode.LengthOfLongestSubstring(s))
//}

//func main() {
//	ransomNote := "a"
//	magazine := "b"
//
//	fmt.Println(leetcode.CanConstruct(ransomNote, magazine))
//
//	ransomNote = "aa"
//	magazine = "ab"
//
//	fmt.Println(leetcode.CanConstruct(ransomNote, magazine))
//
//	ransomNote = "aa"
//	magazine = "aab"
//
//	fmt.Println(leetcode.CanConstruct(ransomNote, magazine))
//}

//func main() {
//s := "egg"
//t := "add"

//fmt.Println(leetcode.IsIsomorphic(s, t))

//s = "foo"
//t = "bar"

//fmt.Println(leetcode.IsIsomorphic(s, t))

//s = "paper"
//t = "title"

//fmt.Println(leetcode.IsIsomorphic(s, t))
//}

//func main() {
//pattern := "abba"
//s := "dog cat cat dog"

//fmt.Println(leetcode.WordPattern(pattern, s))

//pattern = "abba"
//s = "dog cat cat fish"

//fmt.Println(leetcode.WordPattern(pattern, s))

//pattern = "aaaa"
//s = "dog cat cat dog"

//fmt.Println(leetcode.WordPattern(pattern, s))
//}

//func main() {
//	s := "anagram"
//	t := "nagaram"
//	fmt.Println(leetcode.IsAnagram(s, t))
//
//	s = "rat"
//	t = "car"
//	fmt.Println(leetcode.IsAnagram(s, t))
//}

//func main() {
//	s := "()"
//	fmt.Println(leetcode.IsValid(s))
//
//	s = "()[]{}"
//	fmt.Println(leetcode.IsValid(s))
//
//	s = "(]"
//	fmt.Println(leetcode.IsValid(s))
//}

//func main() {
//	s := "/home/"
//	fmt.Println(leetcode.SimplifyPath(s))
//	s = "/../"
//	fmt.Println(leetcode.SimplifyPath(s))
//	s = "/home//foo/"
//	fmt.Println(leetcode.SimplifyPath(s))
//	s = "/a/./b/../../c/"
//	fmt.Println(leetcode.SimplifyPath(s))
//	s = "/home//foo/cc/"
//	fmt.Println(leetcode.SimplifyPath(s))
//}

//func main() {
//minStack := leetcode.Constructor()
//minStack.Push(-2)
//minStack.Push(0)
//minStack.Push(-3)
//fmt.Println(minStack.GetMin())
//minStack.Pop()
//fmt.Println(minStack.Top())
//fmt.Println(minStack.GetMin())
//}

//func main() {
//	tokens := []string{"2", "1", "+", "3", "*"}
//	fmt.Println(leetcode.EvalRPN(tokens))
//	tokens = []string{"4", "13", "5", "/", "+"}
//	fmt.Println(leetcode.EvalRPN(tokens))
//	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
//	fmt.Println(leetcode.EvalRPN(tokens))
//	tokens = []string{"10", "6", "-"}
//	fmt.Println(leetcode.EvalRPN(tokens))
//}

//func main() {
//	s := "1-(     -2)"
//	fmt.Println(leetcode.Calculate(s))
//
//	s = "2147483647"
//	fmt.Println(leetcode.Calculate(s))
//	s = "1 + 1"
//	fmt.Println(leetcode.Calculate(s))
//	s = " 2-1 + 2 "
//	fmt.Println(leetcode.Calculate(s))
//	s = "(1+(4+5+2)-3)+(6+8)"
//	fmt.Println(leetcode.Calculate(s))
//}

//func main() {
//nums := []int{2, 7, 11, 15}
//target := 9
//fmt.Println(leetcode.TwoSum2(nums, target))

//nums = []int{3, 2, 4}
//target = 6
//fmt.Println(leetcode.TwoSum2(nums, target))
//nums = []int{3, 3}
//target = 6
//fmt.Println(leetcode.TwoSum2(nums, target))
//}

//func main() {
//	n := 19
//	fmt.Println(leetcode.IsHappy(n))
//
//	n = 4
//	fmt.Println(leetcode.IsHappy(n))
//}

//func main() {
//	nums := []int{1, 2, 3, 1}
//	k := 3
//	fmt.Println(leetcode.ContainsNearbyDuplicate(nums, k))
//
//	nums = []int{1, 0, 1, 1}
//	k = 1
//	fmt.Println(leetcode.ContainsNearbyDuplicate(nums, k))
//
//	nums = []int{1, 2, 3, 1, 2, 3}
//	k = 2
//	fmt.Println(leetcode.ContainsNearbyDuplicate(nums, k))
//}

//func main() {
//	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
//	fmt.Println(leetcode.GroupAnagrams(strs))
//	strs = []string{""}
//	fmt.Println(leetcode.GroupAnagrams(strs))
//	strs = []string{"a"}
//	fmt.Println(leetcode.GroupAnagrams(strs))
//}

//func main() {
//	nums := []int{100, 4, 200, 1, 3, 2}
//	fmt.Println(leetcode.LongestConsecutive(nums))
//	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
//	fmt.Println(leetcode.LongestConsecutive(nums))
//}

//func main() {
//	fmt.Println(leetcode.HasCycle(&leetcode.ListNode{}))
//}

//func main() {
//	t1 := &leetcode.TreeNode{
//		Val: 4,
//	}
//	t2 := &leetcode.TreeNode{
//		Val: 2,
//	}
//	t3 := &leetcode.TreeNode{
//		Val: 6,
//	}
//	t4 := &leetcode.TreeNode{
//		Val: 1,
//	}
//	t5 := &leetcode.TreeNode{
//		Val: 3,
//	}
//	t1.Left = t2
//	t1.Right = t3
//	t2.Left = t4
//	t2.Right = t5
//
//	fmt.Println(leetcode.GetMinimumDifference(t1))
//}
//

//func main() {
//	numCourses := 5
//	prerequisites := [][]int{
//		{2, 0},
//		{2, 1},
//		{3, 2},
//		{4, 2},
//	}
//	fmt.Println(leetcode.CanFinish(numCourses, prerequisites))
//
//	numCourses = 3
//	prerequisites = [][]int{
//		{1, 0},
//		{2, 1},
//		{0, 2},
//	}
//	fmt.Println(leetcode.CanFinish(numCourses, prerequisites))
//}
//

//func main() {
//	numCourses := 5
//	prerequisites := [][]int{
//		{2, 0},
//		{2, 1},
//		{3, 2},
//		{4, 2},
//	}
//	fmt.Println(leetcode.FindOrder(numCourses, prerequisites))
//
//	numCourses = 3
//	prerequisites = [][]int{
//		{1, 0},
//		{2, 1},
//		{0, 2},
//	}
//	fmt.Println(leetcode.FindOrder(numCourses, prerequisites))
//}

//func main() {
//	board := [][]int{
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, 35, -1, -1, 13, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, 15, -1, -1, -1, -1},
//	}
//	fmt.Println(leetcode.SnakesAndLadders(board))
//
//	board = [][]int{
//		{-1, -1},
//		{-1, 3},
//	}
//	fmt.Println(leetcode.SnakesAndLadders(board))
//
//	board = [][]int{
//		{1, 1, -1},
//		{1, 1, 1},
//		{-1, 1, 1},
//	}
//	fmt.Println(leetcode.SnakesAndLadders(board))
//}

//func main() {
//	fmt.Println(leetcode.Combine(4, 2))
//}

//func main() {
//	nums := []int{1, 2, 3}
//	fmt.Println(leetcode.Permute(nums))
//
//	nums = []int{0, 1}
//	fmt.Println(leetcode.Permute(nums))
//
//	nums = []int{1}
//	fmt.Println(leetcode.Permute(nums))
//}

//func main() {
//	candidates := []int{2, 3, 6, 7}
//	target := 7
//	fmt.Println(leetcode.CombinationSum(candidates, target))
//
//	candidates = []int{2, 3, 5}
//	target = 8
//	fmt.Println(leetcode.CombinationSum(candidates, target))
//
//	candidates = []int{2}
//	target = 1
//	fmt.Println(leetcode.CombinationSum(candidates, target))
//
//}

//func main() {
//	s := "III"
//	fmt.Println(leetcode.RomanToInt(s))
//
//	s = "IV"
//	fmt.Println(leetcode.RomanToInt(s))
//
//	s = "IX"
//	fmt.Println(leetcode.RomanToInt(s))
//
//	s = "LVIII"
//	fmt.Println(leetcode.RomanToInt(s))
//
//	s = "MCMXCIV"
//	fmt.Println(leetcode.RomanToInt(s))
//}

//func main() {
//	num := 3
//	fmt.Println(leetcode.IntToRoman(num))
//
//	num = 4
//	fmt.Println(leetcode.IntToRoman(num))
//
//	num = 9
//	fmt.Println(leetcode.IntToRoman(num))
//
//	num = 58
//	fmt.Println(leetcode.IntToRoman(num))
//
//	num = 1994
//	fmt.Println(leetcode.IntToRoman(num))
//}

//func main() {
//	s := "Hello World"
//	fmt.Println(leetcode.LengthOfLastWord(s))
//	s = "   fly me   to   the moon  "
//	fmt.Println(leetcode.LengthOfLastWord(s))
//	s = "luffy is still joyboy"
//	fmt.Println(leetcode.LengthOfLastWord(s))
//}

//func main() {
//	strs := []string{"flower", "flow", "flight"}
//	fmt.Println(leetcode.LongestCommonPrefix(strs))
//
//	strs = []string{"dog", "racecar", "car"}
//	fmt.Println(leetcode.LongestCommonPrefix(strs))
//}

//func main() {
//t1 := &leetcode.ListNode{
//Val: 4,
//}
//t2 := &leetcode.ListNode{
//Val: 2,
//}
//t3 := &leetcode.ListNode{
//Val: 1,
//}
//t4 := &leetcode.ListNode{
//Val: 3,
//}

//head := t1
//t1.Next = t2
//t2.Next = t3
//t3.Next = t4

//head = leetcode.SortList(head)
//for head != nil {
//fmt.Println(head.Val)
//head = head.Next
//}

//t1 = &leetcode.ListNode{
//Val: -1,
//}
//t2 = &leetcode.ListNode{
//Val: 5,
//}
//t3 = &leetcode.ListNode{
//Val: 3,
//}
//t4 = &leetcode.ListNode{
//Val: 4,
//}
//t5 := &leetcode.ListNode{
//Val: 0,
//}

//head = t1
//t1.Next = t2
//t2.Next = t3
//t3.Next = t4
//t4.Next = t5

//head = leetcode.SortList(head)
//for head != nil {
//fmt.Println(head.Val)
//head = head.Next
//}
//}

//func main() {
//	t1 := &leetcode.ListNode{
//		Val: 1,
//	}
//	t2 := &leetcode.ListNode{
//		Val: 4,
//	}
//	t3 := &leetcode.ListNode{
//		Val: 5,
//	}
//	t4 := &leetcode.ListNode{
//		Val: 1,
//	}
//	t5 := &leetcode.ListNode{
//		Val: 3,
//	}
//	t6 := &leetcode.ListNode{
//		Val: 4,
//	}
//	t7 := &leetcode.ListNode{
//		Val: 2,
//	}
//	t8 := &leetcode.ListNode{
//		Val: 6,
//	}
//	t1.Next = t2
//	t2.Next = t3
//
//	t4.Next = t5
//	t5.Next = t6
//
//	t7.Next = t8
//
//	lists := []*leetcode.ListNode{
//		t1,
//		t4,
//		t7,
//	}
//
//	head := leetcode.MergeKLists(lists)
//	for head != nil {
//		fmt.Println(head.Val)
//		head = head.Next
//	}
//
//	lists = []*leetcode.ListNode{}
//	head = leetcode.MergeKLists(lists)
//	for head != nil {
//		fmt.Println(head.Val)
//		head = head.Next
//	}
//
//	lists = []*leetcode.ListNode{
//		nil,
//	}
//	head = leetcode.MergeKLists(lists)
//	for head != nil {
//		fmt.Println(head.Val)
//		head = head.Next
//	}
//
//}

//func main() {
//	grid := [][]int{
//		{0, 1},
//		{1, 0},
//	}
//	fmt.Println(leetcode.Construct(grid))
//
//	grid = [][]int{
//		{1, 1, 1, 1, 0, 0, 0, 0},
//		{1, 1, 1, 1, 0, 0, 0, 0},
//		{1, 1, 1, 1, 1, 1, 1, 1},
//		{1, 1, 1, 1, 1, 1, 1, 1},
//		{1, 1, 1, 1, 0, 0, 0, 0},
//		{1, 1, 1, 1, 0, 0, 0, 0},
//		{1, 1, 1, 1, 0, 0, 0, 0},
//		{1, 1, 1, 1, 0, 0, 0, 0},
//	}
//	fmt.Println(leetcode.Construct(grid))
//}

//func main() {
//	nums := []int{1, 3, 5, 6}
//	target := 5
//	fmt.Println(leetcode.SearchInsert(nums, target))
//	nums = []int{1, 3, 5, 6}
//	target = 2
//	fmt.Println(leetcode.SearchInsert(nums, target))
//	nums = []int{1, 3, 5, 6}
//	target = 7
//	fmt.Println(leetcode.SearchInsert(nums, target))
//}

//func main() {
//	matrix := [][]int{
//		{1, 3, 5, 7},
//		{10, 11, 16, 20},
//		{23, 30, 34, 60},
//	}
//	target := 3
//	fmt.Println(leetcode.SearchMatrix(matrix, target))
//
//	matrix = [][]int{
//		{1, 3, 5, 7},
//		{10, 11, 16, 20},
//		{23, 30, 34, 60},
//	}
//	target = 13
//	fmt.Println(leetcode.SearchMatrix(matrix, target))
//
//}

//func main() {
//nums := []int{1, 2, 3, 1}
//fmt.Println(leetcode.FindPeakElement(nums))

//nums = []int{1, 2, 1, 3, 5, 6, 4}
//fmt.Println(leetcode.FindPeakElement(nums))

//nums = []int{1, 2}
//fmt.Println(leetcode.FindPeakElement(nums))
//}

//func main() {
//	nums := []int{4, 5, 6, 7, 0, 1, 2}
//	target := 0
//	fmt.Println(leetcode.Search(nums, target))
//	nums = []int{4, 5, 6, 7, 0, 1, 2}
//	target = 3
//	fmt.Println(leetcode.Search(nums, target))
//	nums = []int{1}
//	target = 0
//	fmt.Println(leetcode.Search(nums, target))
//}

//func main() {
//	nums := []int{5, 7, 7, 8, 8, 10}
//	target := 8
//	fmt.Println(leetcode.SearchRange(nums, target))
//
//	nums = []int{5, 7, 7, 8, 8, 10}
//	target = 6
//	fmt.Println(leetcode.SearchRange(nums, target))
//	nums = []int{}
//	target = 0
//	fmt.Println(leetcode.SearchRange(nums, target))
//}

//func main() {
//	digits := []int{2, 9, 9}
//	fmt.Println(leetcode.PlusOne(digits))
//	digits = []int{1, 2, 3}
//	fmt.Println(leetcode.PlusOne(digits))
//	digits = []int{4, 3, 2, 1}
//	fmt.Println(leetcode.PlusOne(digits))
//	digits = []int{0}
//	fmt.Println(leetcode.PlusOne(digits))
//	digits = []int{9, 9, 9}
//	fmt.Println(leetcode.PlusOne(digits))
//	digits = []int{2, 9, 9}
//	fmt.Println(leetcode.PlusOne(digits))
//}

//func main() {
//	a := "11"
//	b := "1"
//	fmt.Println(leetcode.AddBinary(a, b))
//
//	a = "1010"
//	b = "1011"
//	fmt.Println(leetcode.AddBinary(a, b))
//}

//func main() {
//	n := uint32(43261596)
//	fmt.Println(leetcode.ReverseBits(n))
//	//964176192
//
//	n = uint32(4294967293)
//	fmt.Println(leetcode.ReverseBits(n))
//	//3221225471
//}
//

//func main() {
//	n := uint32(11)
//	fmt.Println(leetcode.HammingWeight(n))
//	//3
//}

//func main() {
//	nums := []int{2, 2, 1}
//	fmt.Println(leetcode.SingleNumber(nums))
//	nums = []int{4, 1, 2, 1, 2}
//	fmt.Println(leetcode.SingleNumber(nums))
//	nums = []int{1}
//	fmt.Println(leetcode.SingleNumber(nums))
//}

//func main() {
//nums := []int{2, 2, 3, 2}
//fmt.Println(leetcode.SingleNumber2(nums))
//nums = []int{0, 1, 0, 1, 0, 1, 99}
//fmt.Println(leetcode.SingleNumber2(nums))
//nums = []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}
//fmt.Println(leetcode.SingleNumber2(nums))
//}

//func main() {
//	left := 5
//	right := 7
//	fmt.Println(leetcode.RangeBitwiseAnd(left, right))
//	left = 0
//	right = 0
//	fmt.Println(leetcode.RangeBitwiseAnd(left, right))
//	left = 1
//	right = 2147483647
//	fmt.Println(leetcode.RangeBitwiseAnd(left, right))
//}

//func main() {
//	n := 3
//	fmt.Println(leetcode.TrailingZeroes(n))
//	n = 5
//	fmt.Println(leetcode.TrailingZeroes(n))
//	n = 0
//	fmt.Println(leetcode.TrailingZeroes(n))
//	n = 10
//	fmt.Println(leetcode.TrailingZeroes(n))
//	n = 30
//	fmt.Println(leetcode.TrailingZeroes(n))
//}

//func main() {
//	x := 4
//	fmt.Println(leetcode.MySqrt(x))
//	x = 8
//	fmt.Println(leetcode.MySqrt(x))
//	x = 9
//	fmt.Println(leetcode.MySqrt(x))
//	x = 11
//	fmt.Println(leetcode.MySqrt(x))
//}

//func main() {
//haystack := "sadbutsad"
//needle := "sad"
//fmt.Println(leetcode.StrStr(haystack, needle))

//haystack = "leetcode"
//needle = "leeto"
//fmt.Println(leetcode.StrStr(haystack, needle))
//}

//func main() {
//nums := []int{3, 2, 1, 5, 6, 4}
//k := 2
//fmt.Println(leetcode.FindKthLargest(nums, k))
//nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
//k = 4
//fmt.Println(leetcode.FindKthLargest(nums, k))

//nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
//leetcode.HeapSort(nums)
//fmt.Println(nums)
//}

//func main() {
//elems := []leetcode.Elem{
//{
//Num1: 1,
//Num2: 2,
//},
//{
//Num1: 2,
//Num2: 3,
//},
//}
//elems[0], elems[1] = elems[1], elems[0]

//nums1 := []int{-10, -4, 0, 0, 6, 6, 6, 7, 10, 22, 100}
//nums2 := []int{3, 5, 6, 7, 8, 100}
//k := 20
//fmt.Println(leetcode.KSmallestPairs(nums1, nums2, k))

//nums1 = []int{0, 0, 0, 0, 0, 2, 2, 2, 2}
//nums2 = []int{-3, 22, 35, 56, 76}
//k = 22
//fmt.Println(leetcode.KSmallestPairs(nums1, nums2, k))

//nums1 = []int{1}
//nums2 = []int{3, 5, 6, 7, 8, 100}
//k = 4
//fmt.Println(leetcode.KSmallestPairs(nums1, nums2, k))

//nums1 = []int{1, 7, 11}
//nums2 = []int{2, 4, 6}
//k = 3
//fmt.Println(leetcode.KSmallestPairs(nums1, nums2, k))

//nums1 = []int{1, 1, 2}
//nums2 = []int{1, 2, 3}
//k = 2
//fmt.Println(leetcode.KSmallestPairs(nums1, nums2, k))

//nums1 = []int{1, 2}
//nums2 = []int{3}
//k = 3

//fmt.Println(leetcode.KSmallestPairs(nums1, nums2, k))
//}

//func main() {
//	nums := []int{2, 3, 1, 1, 4}
//	fmt.Println(leetcode.Jump(nums))
//
//	nums = []int{2, 3, 0, 1, 4}
//	fmt.Println(leetcode.Jump(nums))
//}

//func main() {
//	citations := []int{3, 0, 6, 1, 5}
//	fmt.Println(leetcode.HIndex(citations))
//	citations = []int{1, 3, 1}
//	fmt.Println(leetcode.HIndex(citations))
//	citations = []int{1}
//	fmt.Println(leetcode.HIndex(citations))
//}

//func main() {
//nums := []int{1, 2, 3, 4}
//fmt.Println(leetcode.ProductExceptSelf(nums))
//nums = []int{-1, 1, 0, -3, 3}
//fmt.Println(leetcode.ProductExceptSelf(nums))
//}

//func main() {
//	gas := []int{1, 2, 3, 4, 5}
//	cost := []int{3, 4, 5, 1, 2}
//	fmt.Println(leetcode.CanCompleteCircuit(gas, cost))
//
//	gas = []int{2, 3, 4}
//	cost = []int{3, 4, 3}
//	fmt.Println(leetcode.CanCompleteCircuit(gas, cost))
//}

//func main() {
//s := "the sky is blue"
//fmt.Println(leetcode.ReverseWords(s))
//s = "  hello world  "
//fmt.Println(leetcode.ReverseWords(s))
//s = "a good   example"
//fmt.Println(leetcode.ReverseWords(s))
//}

//func main() {
//	trie := leetcode.ConstructorTrie()
//	trie.Insert("apple")
//	fmt.Println(trie.Search("apple"))   // 返回 True
//	fmt.Println(trie.Search("app"))     // 返回 False
//	fmt.Println(trie.StartsWith("app")) // 返回 True
//	trie.Insert("app")
//	fmt.Println(trie.Search("app")) // 返回 True
//}

//func main() {
//	wordDictionary := leetcode.ConstructorWordDictionary()
//	wordDictionary.AddWord("at")
//	wordDictionary.AddWord("and")
//	wordDictionary.AddWord("an")
//	wordDictionary.AddWord("add")
//	fmt.Println(wordDictionary.Search("a"))   // 返回
//	fmt.Println(wordDictionary.Search(".at")) // 返回
//	wordDictionary.AddWord("bat")
//	fmt.Println(wordDictionary.Search(".at"))  // 返回
//	fmt.Println(wordDictionary.Search("an."))  // 返回
//	fmt.Println(wordDictionary.Search("a.d.")) // 返回
//	fmt.Println(wordDictionary.Search("b."))   // 返回
//	fmt.Println(wordDictionary.Search("a.d"))  // 返回
//	fmt.Println(wordDictionary.Search("."))    // 返回
//
//	wordDictionary = leetcode.ConstructorWordDictionary()
//	wordDictionary.AddWord("bad")
//	wordDictionary.AddWord("dad")
//	wordDictionary.AddWord("mad")
//	fmt.Println(wordDictionary.Search("pad")) // 返回 False
//	fmt.Println(wordDictionary.Search("bad")) // 返回 True
//	fmt.Println(wordDictionary.Search(".ad")) // 返回 True
//	fmt.Println(wordDictionary.Search("b..")) // 返回 True
//}

//func main() {
//	board := [][]byte{
//		{'o', 'a', 'a', 'n'},
//		{'e', 't', 'a', 'e'},
//		{'i', 'h', 'k', 'r'},
//		{'i', 'f', 'l', 'v'},
//	}
//	words := []string{"oath", "pea", "eat", "rain"}
//	fmt.Println(leetcode.FindWords(board, words))
//
//	board = [][]byte{
//		{'a', 'b'},
//		{'c', 'd'},
//	}
//	words = []string{"abcd"}
//	fmt.Println(leetcode.FindWords(board, words))
//}

//func main() {
//	lru := leetcode.Constructor(3)
//	lru.Put(1, 1)
//	lru.Put(2, 2)
//	lru.Put(3, 3)
//	lru.Put(4, 4)
//	fmt.Println(lru.Get(4))
//	fmt.Println(lru.Get(3))
//	fmt.Println(lru.Get(2))
//	fmt.Println(lru.Get(1))
//	lru.Put(5, 5)
//	fmt.Println(lru.Get(1))
//	fmt.Println(lru.Get(2))
//	fmt.Println(lru.Get(3))
//	fmt.Println(lru.Get(4))
//	fmt.Println(lru.Get(5))
//}

//func main() {
//board := [][]byte{
//{'A', 'B', 'C', 'E'},
//{'S', 'F', 'C', 'S'},
//{'A', 'D', 'E', 'E'},
//}

//word := "ABCCED"
//fmt.Println(leetcode.Exist(board, word))

//board = [][]byte{
//{'A', 'B', 'C', 'E'},
//{'S', 'F', 'C', 'S'},
//{'A', 'D', 'E', 'E'},
//}
//word = "SEE"
//fmt.Println(leetcode.Exist(board, word))

//board = [][]byte{
//{'A', 'B', 'C', 'E'},
//{'S', 'F', 'C', 'S'},
//{'A', 'D', 'E', 'E'},
//}
//word = "ABCB"
//fmt.Println(leetcode.Exist(board, word))
//}

//func main() {
//	s := "a"
//	fmt.Println(leetcode.LongestPalindrome(s))
//	s = "cb"
//	fmt.Println(leetcode.LongestPalindrome(s))
//	s = "bb"
//	fmt.Println(leetcode.LongestPalindrome(s))
//	s = "babad"
//	fmt.Println(leetcode.LongestPalindrome(s))
//	s = "cbbd"
//	fmt.Println(leetcode.LongestPalindrome(s))
//}

//func main() {
//	s1 := "aabcc"
//	s2 := "dbbca"
//	s3 := "aadbbcbcac"
//
//	fmt.Println(leetcode.IsInterleave(s1, s2, s3))
//
//	s1 = "aabcc"
//	s2 = "dbbca"
//	s3 = "aadbbbaccc"
//
//	fmt.Println(leetcode.IsInterleave(s1, s2, s3))
//}

//func main() {
//	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
//	fmt.Println(leetcode.LengthOfLIS(nums))
//	nums = []int{0, 1, 0, 3, 2, 3}
//	fmt.Println(leetcode.LengthOfLIS(nums))
//	nums = []int{7, 7, 7, 7, 7, 7, 7}
//	fmt.Println(leetcode.LengthOfLIS(nums))
//}

//func main() {
//	//set := leetcode.ConstructRandomizedSet()
//	//fmt.Println(set.Insert(1))
//	//fmt.Println(set.Remove(2))
//	//fmt.Println(set.Insert(2))
//	//fmt.Println(set.GetRandom())
//	//fmt.Println(set.Remove(1))
//	//fmt.Println(set.Insert(2))
//	//fmt.Println(set.GetRandom())
//
//	//set = leetcode.ConstructRandomizedSet()
//	//fmt.Println(set.Remove(0))
//	//fmt.Println(set.Remove(0))
//	//fmt.Println(set.Insert(0))
//	//fmt.Println(set.GetRandom())
//	//fmt.Println(set.Remove(0))
//	//fmt.Println(set.Insert(0))
//
//	set := leetcode.ConstructRandomizedSet()
//	fmt.Println(set.Insert(3))
//	fmt.Println(set.Remove(3))
//	fmt.Println(set.Remove(0))
//	fmt.Println(set.Insert(3))
//	fmt.Println(set.GetRandom())
//	fmt.Println(set.Remove(0))
//}

//func main() {
//	s := "AB"
//	numRows := 1
//	fmt.Println(leetcode.Convert(s, numRows))
//
//	s = "PAYPALISHIRING"
//	numRows = 3
//	fmt.Println(leetcode.Convert(s, numRows))
//
//	s = "PAYPALISHIRING"
//	numRows = 4
//	fmt.Println(leetcode.Convert(s, numRows))
//}

//func main() {
//t1 := &leetcode.ListNode{
//Val: 1,
//}
//t2 := &leetcode.ListNode{
//Val: 2,
//}
//t3 := &leetcode.ListNode{
//Val: 3,
//}
//t4 := &leetcode.ListNode{
//Val: 4,
//}
//t5 := &leetcode.ListNode{
//Val: 5,
//}
//t1.Next = t2
//t2.Next = t3
//t3.Next = t4
//t4.Next = t5
//left := 2
//right := 4
//head := leetcode.ReverseBetween(t1, left, right)
//for head != nil {
//fmt.Println(head.Val)
//head = head.Next
//}

//fmt.Println("case2 ..............")
//t1 = &leetcode.ListNode{
//Val: 5,
//}
//left = 1
//right = 1
//head = leetcode.ReverseBetween(t1, left, right)
//for head != nil {
//fmt.Println(head.Val)
//head = head.Next
//}

//fmt.Println("case3 ..............")
//t1 = &leetcode.ListNode{
//Val: 3,
//}
//t2 = &leetcode.ListNode{
//Val: 5,
//}
//t1.Next = t2
//left = 1
//right = 2
//head = leetcode.ReverseBetween(t1, left, right)
//for head != nil {
//fmt.Println(head.Val)
//head = head.Next
//}

//}

//func main() {
//	numbers := []int{1, 2, 3, 4, 5}
//	//helper := func(numbers []int) {
//	//numbers[0] = 5
//	//for i := 0; i < len(numbers); i++ {
//	//fmt.Println(numbers[i])
//	//}
//	//}
//
//	//helper(numbers[:3])
//	numbers2 := numbers[2:]
//	numbers2[0] = 9
//	for i := 0; i < len(numbers); i++ {
//		fmt.Println(numbers[i])
//	}
//}

//func main() {
//t1 := &leetcode.TreeNode{Val: 7}
//t2 := &leetcode.TreeNode{Val: 3}
//t3 := &leetcode.TreeNode{Val: 15}
//t4 := &leetcode.TreeNode{Val: 9}
//t5 := &leetcode.TreeNode{Val: 20}
//t1.Left = t2
//t1.Right = t3
//t3.Left = t4
//t3.Right = t5

//bst := leetcode.ConstructorBSTIterator(t1)
//fmt.Println(bst.Next())
//fmt.Println(bst.Next())
//fmt.Println(bst.HasNext())
//fmt.Println(bst.Next())
//fmt.Println(bst.HasNext())
//fmt.Println(bst.Next())
//fmt.Println(bst.HasNext())
//fmt.Println(bst.Next())
//fmt.Println(bst.HasNext())

//bst = leetcode.ConstructorBSTIterator(t5)
//fmt.Println(bst.HasNext())
//fmt.Println(bst.Next())
//}

//func main() {
//board := [][]byte{
//{'1', '1'},
//}
//checker := func(i, j int, sign []bool) bool {
//c := board[i][j]
//if c == '.' {
//return false
//}
//if sign[c-'0'] {
//return true
//}

//sign[c-'0'] = true
//return false
//}

//sign := make([]bool, 10)
//fmt.Println(checker(0, 0, sign))
//fmt.Println(checker(0, 1, sign))

//}

//func main() {
//	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
//	leetcode.SetZeroes(matrix)
//}

//func main() {
//	intervals := [][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}
//	fmt.Println(leetcode.MergeIntervals(intervals))
//
//}

//func main() {
//intervals := [][]int{
//{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
//}
//newInterval := []int{4, 8}
//fmt.Println(leetcode.InsertInterval(intervals, newInterval))

//}

//func main() {
//points := [][]int{
//{10, 16}, {2, 8}, {1, 6}, {7, 12},
//}
//fmt.Println(leetcode.FindMinArrowShots(points))

//points = [][]int{
//{1, 2}, {3, 4}, {5, 6}, {7, 8},
//}
//fmt.Println(leetcode.FindMinArrowShots(points))

//points = [][]int{
//{1, 2}, {2, 3}, {3, 4}, {4, 5},
//}
//fmt.Println(leetcode.FindMinArrowShots(points))

//points = [][]int{
//{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7},
//}
//fmt.Println(leetcode.FindMinArrowShots(points))
//}

//func main() {
//	n1 := &graph.Node{
//		Val: 1,
//	}
//	n2 := &graph.Node{
//		Val: 2,
//	}
//	n3 := &graph.Node{
//		Val: 3,
//	}
//	n4 := &graph.Node{
//		Val: 4,
//	}
//
//	n1.Neighbors = []*graph.Node{n2, n4}
//	n2.Neighbors = []*graph.Node{n1, n3}
//	n3.Neighbors = []*graph.Node{n2, n4}
//	n4.Neighbors = []*graph.Node{n1, n3}
//
//	graph.CloneGraph(n1)
//}

//func main() {
//matrix := [][]byte{
//{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'},
//}

//fmt.Println(leetcode.MaximalSquare(matrix))

//matrix = [][]byte{
//{'0', '1'}, {'1', '0'},
//}
//fmt.Println(leetcode.MaximalSquare(matrix))
//}

//func main() {
//	x := 2.00000
//	n := 10
//	fmt.Println(leetcode.MyPow(x, n))
//	x = 2.10000
//	n = 3
//	fmt.Println(leetcode.MyPow(x, n))
//	x = 2.00000
//	n = -2
//	fmt.Println(leetcode.MyPow(x, n))
//}

//func main() {
//n := 4
//fmt.Println(leetcode.TotalNQueens(n))
//n = 1
//fmt.Println(leetcode.TotalNQueens(n))
//}

//func main() {
//nums := []int{3, 4, 5, 1, 2}
//fmt.Println(leetcode.FindMin(nums))

//nums = []int{4, 5, 6, 7, 0, 1, 2}
//fmt.Println(leetcode.FindMin(nums))

//nums = []int{11, 13, 15, 17}
//fmt.Println(leetcode.FindMin(nums))
//}

//func main() {
//nums := []int{1, -2, 3, -2}
//fmt.Println(leetcode.MaxSubarraySumCircular(nums))

//nums = []int{5, -3, 5}
//fmt.Println(leetcode.MaxSubarraySumCircular(nums))

//nums = []int{3, -2, 2, -3}
//fmt.Println(leetcode.MaxSubarraySumCircular(nums))

//nums = []int{-3, -2, -3}
//fmt.Println(leetcode.MaxSubarraySumCircular(nums))
//}

//func main() {
//	data := [][]int{
//		{0, 4}, {1, 5}, {3, 4}, {9, 13},
//	}
//	fmt.Println(leetcode.MergeData(data))
//}

//func main() {
//var sMp sync.Map
//sMp.Store("key1", "value1")
//sMp.Store("key2", "value2")
//sMp.Store("key3", "value3")

//fmt.Println(sMp.Load("key3"))
//sMp.Range(func(key, value any) bool {
//fmt.Printf("%v\n", key)
//return key == "key1"
//})
//}

//func main() {
//	ratings := []int{1, 0, 2}
//	fmt.Println(leetcode.Candy(ratings))
//	ratings = []int{1, 2, 2}
//	fmt.Println(leetcode.Candy(ratings))
//}

//func main() {
//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//fmt.Println(leetcode.Trap(height))

//height = []int{4, 2, 0, 3, 2, 5}
//fmt.Println(leetcode.Trap(height))
//}

//func main() {
//words := []string{"This", "is", "an", "example", "of", "text", "justification."}
//maxWidth := 16
//fmt.Println(leetcode.FullJustify(words, maxWidth))

//words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
//maxWidth = 16
//fmt.Println(leetcode.FullJustify(words, maxWidth))

//words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
//maxWidth = 20
//fmt.Println(leetcode.FullJustify(words, maxWidth))
//}

//func main() {
////data := "()()"
//data := "(())"
//fmt.Println(leetcode.MaxLen(data))
////data = "(())"
////fmt.Println(leetcode.MaxLen(data))
////data = "()(()"
////fmt.Println(leetcode.MaxLen(data))
////data = "()(()()(()"
////fmt.Println(leetcode.MaxLen(data))
//}

//func main() {
//s := "())"
//fmt.Println(leetcode.LongestValidParentheses(s))
//s = "(()"
//fmt.Println(leetcode.LongestValidParentheses(s))
//s = ")()())"
//fmt.Println(leetcode.LongestValidParentheses(s))
//s = ""
//fmt.Println(leetcode.LongestValidParentheses(s))
//}

//func main() {
//	s := "wordgoodgoodgoodbestword"
//	words := []string{"word", "good", "best", "good"}
//	fmt.Println(leetcode.FindSubstring(s, words))
//	//s := "ababababab"
//	//words := []string{"ababa", "babab"}
//	//fmt.Println(leetcode.FindSubstring(s, words))
//	//s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
//	//words := []string{"fooo", "barr", "wing", "ding", "wing"}
//	//fmt.Println(leetcode.FindSubstring(s, words))
//	//s = "barfoothefoobarman"
//	//words = []string{"foo", "bar"}
//	//fmt.Println(leetcode.FindSubstring(s, words))
//	//s = "wordgoodgoodgoodbestword"
//	//words = []string{"word", "good", "best", "word"}
//	//fmt.Println(leetcode.FindSubstring(s, words))
//	//s = "barfoofoobarthefoobarman"
//	//words = []string{"bar", "foo", "the"}
//	//fmt.Println(leetcode.FindSubstring(s, words))
//}

//func main() {
//s := "ADOBECODEBANC"
//t := "ABC"
//fmt.Println(leetcode.MinWindow(s, t))
//s = "a"
//t = "a"
//fmt.Println(leetcode.MinWindow(s, t))
//s = "a"
//t = "aa"
//fmt.Println(leetcode.MinWindow(s, t))
//}

//func main() {
//nums := []int{2, 2, 1}
//fmt.Println(leetcode.SingleNumber4(nums))

//nums = []int{4, 1, 2, 1, 2}
//fmt.Println(leetcode.SingleNumber4(nums))
//}

//func main() {
//s := "42"
//fmt.Println(leetcode.MyAtoi1111(s))

//s = "-42"
//fmt.Println(leetcode.MyAtoi1111(s))

//s = "4193 with words"
//fmt.Println(leetcode.MyAtoi1111(s))

//s = "abdc dbc -4193 with words"
//fmt.Println(leetcode.MyAtoi1111(s))
//s = "2147483648"
//fmt.Println(leetcode.MyAtoi1111(s))
//s = "-2147483649"
//fmt.Println(leetcode.MyAtoi1111(s))
//fmt.Println(math.MaxInt32)
//fmt.Println(math.MinInt32)
//}

///func main() {
///k := 111
///w := 12
///profits := []int{319, 178, 37, 756, 152, 763, 481, 1055, 594, 825, 759, 494, 1087, 696, 717, 358, 1091, 1145, 1043, 245, 779, 957, 3, 1060, 1141, 926, 226, 657, 869, 740, 1170, 746, 889, 416, 634, 209, 1189, 1076, 819, 1144, 40, 807, 561, 400, 283, 133, 186, 839, 1043, 491, 936, 559, 70, 9, 854, 1172, 9, 739, 871, 436, 1087, 1029, 66, 561, 798, 1141, 701, 1020, 1072, 92, 636, 509, 392, 77, 84, 202, 843, 563, 329, 906, 101, 997, 1162, 105, 876, 460, 285, 446, 753, 516, 60, 932, 323, 659, 9, 639, 1041, 861, 1199, 343, 792, 1114, 536, 405, 542, 805, 929, 589, 538, 410, 143, 114, 24, 1064, 682, 536, 1016, 1141, 843, 472, 817, 196, 673, 189, 915, 196, 755, 203, 956, 283, 833, 836, 22, 899, 232, 655, 572, 207, 819, 639, 1024, 887, 407, 385, 251, 896, 165, 290, 193, 357, 870, 678, 411, 697, 998, 344, 628, 866, 1004, 894, 547, 574, 36, 847, 691, 601, 896, 363, 21, 1133, 115, 658, 855, 1042, 2, 189, 230, 215, 747, 971, 581, 809, 832, 311, 35, 639, 1127, 701, 167, 672, 763, 997, 1061, 1170, 289, 806, 91, 198, 720, 67, 273, 863, 437, 152, 671, 635, 968, 107, 58, 453, 392, 236, 875, 719, 44, 850, 176, 580, 937, 227, 125, 851, 1188, 915, 747, 1175, 69, 743, 250, 275, 290, 787, 780, 516, 775, 279, 256, 649, 83, 1180, 372, 655, 813, 750, 381, 807, 988, 977, 430, 282, 1067, 551, 726, 756, 372, 17, 1016, 550, 1128, 223, 174, 85, 1039, 1, 27, 178, 822, 243, 440, 79, 772, 211, 640, 454, 1141, 654, 351, 54, 794, 711, 157, 492, 307, 1156, 937, 212, 461, 417, 741, 1104, 642, 1147, 833, 1194, 1066, 1070, 876, 1045, 341, 854, 1014, 787, 518, 404, 1068, 385, 756, 998, 298, 439, 917, 183, 71, 363, 38, 828, 372, 1071, 195, 481, 1153, 871, 712, 803, 849, 145, 1144, 410, 420, 561, 678, 241, 498, 548, 436, 871, 810, 180, 635, 119, 665, 551, 1132, 807, 408, 685, 489, 274, 459, 28, 987, 935, 310, 284, 1163, 924, 812, 267, 1123, 782, 1065, 1075, 1199, 1047, 289, 280, 1044, 931, 1056, 625, 672, 150, 41, 1084, 998, 151, 483, 226, 548, 277, 1187, 1010, 262, 1048, 493, 1084, 845, 1127, 1009, 773, 129, 547, 798, 118, 559, 463, 972, 666, 1155, 519, 381, 405, 819, 201, 790, 729, 1100, 614, 691, 3, 386, 206, 514, 387, 612, 1073, 158, 67, 1116, 164, 496, 1056, 1159, 348, 530, 28, 1177, 774, 1139, 559, 563, 310, 916, 897, 36, 1060, 173, 952, 123, 635, 911, 711, 304, 611, 481, 645, 532, 1033, 385, 402, 797, 117, 307, 969, 1147, 1127, 434, 1099, 1043, 104, 74, 257, 778, 934, 901, 106, 84, 525, 138, 698, 877, 1009, 487, 450, 649, 736, 802, 1032, 456, 757, 10, 2, 720, 1155, 234, 844, 880, 1053, 1134, 821, 1094, 970, 1095, 976, 101, 1163, 104, 38, 198, 350, 896, 345, 867, 1129, 1064, 403, 228, 1103, 416, 561, 722, 915, 1161, 252, 45, 632, 1163, 232, 174, 964, 876, 1126, 479, 432, 984, 815, 544, 1005, 620, 497, 1021, 763, 169, 1028, 706, 1112, 150, 368, 483, 251, 721, 492, 686, 1105, 38, 1140, 1085, 153, 366, 428, 296, 28, 648, 873, 21, 236, 393, 581, 1043, 1086, 747, 402, 201, 1196, 416, 788, 318, 257, 815, 735, 1023, 1143, 1100, 922, 1033, 884, 824, 963, 159, 904, 21, 1168, 511, 723, 400, 239, 338, 89, 1099, 572, 674, 54, 891, 426, 277, 91, 504, 302, 616, 468, 506, 917, 491, 744, 1091, 1051, 594, 465, 850, 338, 417, 320, 1160, 364, 480, 855, 86, 305, 225, 833, 637, 760, 1147, 613, 236, 460, 664, 1183, 38, 806, 100, 654, 852, 975, 816, 506, 1092, 275, 6, 229, 972, 53, 554, 370, 63, 651, 701, 304, 1011, 672, 353, 118, 1111, 448, 549, 151, 1087, 909, 420, 1095, 663, 1119, 1124, 729, 578, 733, 861, 1154, 1195, 387, 1182, 850, 189, 463, 1129, 1185, 541, 546, 1159, 314, 137, 479, 59, 841, 514, 548, 496, 642, 341, 151, 999, 1036, 1186, 704, 550, 1039, 723, 779, 584, 382, 371, 712, 176, 665, 180, 440, 697, 1102, 408, 728, 157, 1050, 558, 692, 336, 384, 107, 839, 477, 204, 862, 345, 429, 1144, 1069, 207, 449, 594, 963, 988, 843, 334, 200, 865, 281, 296, 1040, 914, 995, 743, 1134, 1050, 484, 602, 1001, 570, 1052, 1106, 922, 77, 1099, 681, 360, 955, 1184, 669, 411, 760, 46, 599, 815, 231, 86, 694, 469, 1079, 1161, 1105, 519, 1010, 681, 603, 376, 534, 145, 361, 448, 1006, 91, 397, 671, 597, 238, 119, 467, 587, 723, 162, 804, 638, 568, 195, 460, 610, 868, 806, 995, 1178, 490, 406, 933, 232, 17, 37, 2, 114, 1004, 871, 531, 209, 267, 37, 750, 1196, 940, 89, 725, 377, 898, 6, 685, 210, 1127, 1160, 432, 391, 931, 681, 937, 275, 1190, 137, 743, 652, 430, 566, 180, 1192, 718, 253, 133, 998, 1067, 764, 747, 1159, 827, 143, 506, 641, 327, 468, 731, 50, 15, 569, 80, 310, 1086, 1092, 989, 245, 520, 711, 788, 1144, 938, 1152, 1169, 563, 1053, 1182, 331, 838, 112, 361, 1049, 717, 979, 956, 434, 534, 1083, 117, 280, 1104, 293, 1137, 592, 1019, 606, 526, 216, 924, 197, 137, 1041, 2, 910, 274, 1178, 267, 521, 626, 764, 691, 124, 446, 337, 676, 325, 288, 1120, 924, 512, 777, 1063, 979, 86, 677, 1183, 777, 418, 41, 852, 929, 712, 1132, 1030, 339, 943, 475, 851, 340, 894, 1091, 230, 654, 229, 945, 69, 182, 110, 255, 895, 645, 1061, 793, 1016, 807, 440, 330, 519, 73, 108, 660, 209, 1077, 1191, 938, 415, 1162, 579, 258, 14, 273, 561, 834, 371, 134, 1118, 1139, 1163, 667, 979, 483, 436, 42, 593, 139, 846, 875, 571, 808, 76, 713, 1198, 352, 299, 156, 793, 509, 245, 697, 1106, 236, 287, 236, 644, 683}
///capital := []int{487, 13, 943, 31, 661, 656, 690, 175, 1147, 760, 96, 290, 755, 504, 1111, 219, 187, 641, 380, 886, 781, 214, 714, 594, 41, 1154, 908, 977, 1183, 28, 464, 524, 895, 1177, 28, 225, 1180, 206, 387, 25, 166, 207, 394, 418, 771, 719, 153, 836, 970, 589, 114, 67, 1180, 1136, 863, 1144, 21, 9, 957, 861, 981, 849, 1031, 361, 541, 649, 220, 718, 263, 390, 24, 830, 103, 955, 233, 1174, 521, 580, 83, 869, 321, 712, 1011, 130, 297, 835, 406, 543, 849, 681, 337, 355, 867, 863, 552, 883, 155, 762, 982, 129, 605, 125, 1111, 442, 89, 1139, 1084, 870, 183, 1142, 78, 12, 893, 677, 817, 348, 1036, 482, 76, 619, 563, 435, 529, 311, 1148, 629, 786, 288, 112, 995, 854, 844, 1002, 948, 570, 736, 1088, 354, 618, 876, 926, 150, 1108, 412, 707, 233, 137, 775, 751, 1137, 481, 349, 150, 518, 1165, 191, 223, 754, 419, 1025, 817, 1001, 278, 692, 403, 1023, 106, 78, 124, 679, 598, 727, 1026, 966, 564, 726, 1148, 643, 806, 1182, 645, 300, 867, 613, 458, 844, 679, 907, 437, 1012, 902, 18, 843, 190, 43, 705, 818, 957, 146, 1175, 130, 744, 941, 975, 692, 1066, 541, 335, 20, 311, 606, 377, 558, 113, 545, 1115, 228, 29, 2, 1180, 331, 1072, 151, 692, 1156, 347, 293, 1135, 959, 865, 1031, 40, 425, 1191, 1178, 87, 989, 92, 1186, 1072, 105, 1058, 369, 401, 1117, 220, 484, 181, 901, 321, 923, 263, 72, 685, 820, 1123, 736, 942, 37, 419, 631, 545, 761, 227, 230, 25, 636, 1048, 834, 736, 899, 530, 217, 669, 278, 653, 657, 857, 724, 782, 146, 780, 615, 1156, 751, 463, 625, 707, 355, 92, 718, 855, 624, 504, 359, 744, 793, 868, 462, 985, 1087, 517, 886, 83, 221, 383, 601, 709, 683, 1097, 544, 411, 28, 1129, 781, 13, 752, 347, 1150, 1030, 269, 701, 658, 658, 1188, 222, 1160, 480, 953, 203, 132, 17, 723, 927, 911, 448, 977, 1126, 219, 118, 1033, 919, 1041, 712, 930, 963, 772, 264, 523, 479, 735, 919, 72, 1019, 131, 15, 628, 331, 408, 578, 37, 1123, 125, 527, 887, 54, 1043, 259, 654, 557, 872, 505, 700, 1077, 202, 368, 628, 29, 66, 199, 611, 730, 1108, 682, 705, 628, 1001, 705, 21, 266, 428, 1093, 800, 742, 701, 715, 845, 1151, 300, 460, 1134, 708, 443, 706, 819, 296, 199, 671, 452, 626, 386, 569, 999, 326, 1081, 202, 384, 783, 541, 811, 1058, 684, 146, 288, 1149, 480, 0, 658, 547, 773, 588, 758, 189, 489, 257, 436, 362, 417, 180, 671, 141, 657, 864, 808, 1026, 504, 827, 1125, 977, 1161, 699, 360, 689, 520, 796, 147, 746, 978, 833, 82, 102, 254, 736, 88, 525, 1037, 659, 1061, 333, 663, 520, 6, 439, 823, 1151, 395, 188, 32, 513, 34, 805, 1017, 1193, 157, 942, 87, 630, 915, 807, 215, 482, 1132, 23, 1008, 1175, 724, 1070, 339, 1139, 22, 455, 306, 369, 151, 1031, 1119, 846, 1195, 49, 1169, 773, 706, 45, 747, 875, 944, 1161, 1025, 258, 765, 1039, 459, 641, 1169, 211, 894, 556, 253, 831, 1115, 458, 351, 1138, 58, 1169, 1190, 743, 611, 392, 1015, 54, 831, 330, 1138, 681, 1012, 750, 967, 1179, 398, 564, 5, 2, 143, 787, 197, 590, 144, 589, 588, 493, 524, 748, 1123, 707, 585, 641, 293, 871, 331, 1173, 280, 218, 931, 11, 663, 1005, 1118, 555, 1000, 699, 720, 81, 527, 71, 458, 462, 1056, 447, 491, 1068, 1078, 842, 455, 497, 959, 745, 654, 1011, 939, 787, 430, 535, 594, 940, 1176, 656, 207, 815, 1133, 610, 1113, 596, 1018, 961, 840, 555, 582, 1187, 828, 161, 983, 686, 870, 800, 525, 34, 1013, 385, 870, 953, 59, 518, 521, 151, 633, 561, 848, 310, 712, 108, 1148, 446, 480, 983, 110, 442, 182, 119, 463, 909, 690, 48, 1040, 631, 17, 1027, 158, 353, 108, 524, 335, 1046, 514, 1027, 746, 917, 378, 437, 606, 829, 743, 462, 1013, 525, 290, 477, 749, 896, 12, 351, 13, 42, 819, 334, 912, 30, 1041, 815, 307, 1138, 168, 209, 1134, 100, 276, 292, 283, 1074, 123, 561, 857, 92, 879, 58, 706, 532, 75, 49, 385, 342, 887, 646, 301, 54, 483, 589, 1084, 1092, 178, 845, 243, 873, 611, 340, 712, 115, 157, 63, 773, 800, 844, 167, 384, 522, 877, 183, 368, 709, 501, 905, 512, 756, 246, 197, 463, 47, 232, 256, 37, 883, 1048, 774, 431, 943, 868, 111, 163, 336, 648, 313, 858, 536, 416, 680, 951, 320, 499, 199, 234, 482, 529, 676, 26, 1180, 721, 877, 586, 628, 1152, 835, 1145, 447, 763, 750, 188, 1169, 596, 1125, 310, 519, 1165, 488, 1063, 59, 292, 701, 1078, 1088, 663, 512, 172, 477, 187, 215, 1192, 22, 827, 279, 607, 286, 179, 744, 569, 500, 510, 1038, 570, 1159, 520, 1070, 759, 831, 906, 644, 753, 574, 257, 983, 1023, 227, 460, 710, 169, 595, 249, 111, 73, 991, 933, 448, 655, 559, 183, 244, 44, 644, 935, 466, 97, 605, 460, 800, 229, 977, 675, 236, 946, 73, 456, 623, 499, 423, 162, 342, 914, 386, 1082, 407, 954, 1081, 1099, 142, 539, 416, 791, 1195, 1099, 885, 965, 971, 796, 1198, 449, 768, 792, 541, 18, 476, 303, 137, 319, 1008, 343, 733, 128, 641, 709, 175, 691, 227, 307, 1177, 1198, 1075, 747, 944, 1038, 933, 643, 613, 1166, 1153, 120, 288, 1167, 246, 1103, 185, 85, 1008, 1060, 1051, 421, 150, 601, 376, 183, 1028, 146, 297, 515, 688, 886, 1090, 552, 969, 903, 1086, 931, 946, 1099, 546, 17, 851, 908, 1170, 418, 94, 61, 233, 1069, 510, 783, 302, 701, 564, 1195, 57, 1007, 994, 205, 1019, 694, 1020, 137, 1041, 803, 673, 1162, 14, 904, 418, 1076, 514, 79, 944, 242, 491, 867, 934, 40, 1059, 776, 817, 468, 516, 550, 906, 790, 459, 273, 924, 185, 1183, 337, 435, 699, 316, 768}
///fmt.Println(leetcode.FindMaximizedCapital(k, w, profits, capital))

///k = 2
///w = 0
///profits = []int{1, 2, 3}
///capital = []int{0, 1, 1}
///fmt.Println(leetcode.FindMaximizedCapital(k, w, profits, capital))

///k = 10
///w = 0
///profits = []int{1, 2, 3}
///capital = []int{0, 1, 2}
///fmt.Println(leetcode.FindMaximizedCapital(k, w, profits, capital))

///k = 1
///w = 0
///profits = []int{1, 2, 3}
///capital = []int{0, 1, 2}
///fmt.Println(leetcode.FindMaximizedCapital(k, w, profits, capital))
///}

//func main() {
//pattern := "abba"
//str := "北京 杭州 杭州 北京"
//fmt.Println(leetcode.CheckStrPattern(pattern, str))
//pattern = "aabb"
//str = "北京 杭州 杭州 北京"
//fmt.Println(leetcode.CheckStrPattern(pattern, str))
//pattern = "baab"
//str = "北京 杭州 杭州 北京"
//fmt.Println(leetcode.CheckStrPattern(pattern, str))
//}

//func main() {
//mf := heap.Constructor()
//mf.AddNum(1)
//mf.AddNum(2)
//fmt.Println(mf.FindMedian())
//mf.AddNum(3)
//fmt.Println(mf.FindMedian())
//}

//func main() {
//s := "cbaebabacd"
//p := "abc"
//fmt.Println(leetcode.FindAnagrams(s, p))

//s = "abab"
//p = "ab"
//fmt.Println(leetcode.FindAnagrams(s, p))
//}

//func main() {
//	t1 := &link.ListNode{Val: 1}
//	t2 := &link.ListNode{Val: 2}
//	t3 := &link.ListNode{Val: 3}
//	t4 := &link.ListNode{Val: 4}
//	t1.Next = t2
//	t2.Next = t3
//	t3.Next = t4
//	fmt.Println(link.SwapPairs(t1))
//}

//func main() {
//grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
//fmt.Println(graph.OrangesRotting(grid))

//grid = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
//fmt.Println(graph.OrangesRotting(grid))

//grid = [][]int{{0, 2}}
//fmt.Println(graph.OrangesRotting(grid))

//grid = [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}
//fmt.Println(graph.OrangesRotting(grid))
//}

//func main() {
//nums := []int{1, 2, 3}
//fmt.Println(backtrace.Subsets(nums))
//nums = []int{0}
//fmt.Println(backtrace.Subsets(nums))
//}

func main() {
	s := "aabaa"
	fmt.Println(backtrace.Partition(s))
}
