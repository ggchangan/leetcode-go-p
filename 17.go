package main

import (
	"strings"
)

var words []string
var digitLetters = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func combination(l int, n int, digits []string, word string) {
	if l == n {
		words = append(words, word)
		return
	}

	for _, c := range digitLetters[digits[l]] {
		newWord := word + string(c)
		combination(l+1, n, digits, newWord)
	}
}

// vscode is right, however this cannot pass in leetcode
func letterCombinations(digits string) []string {
	words = make([]string, 0)
	digitLetters2 := strings.Split(digits, ",")
	combination(0, len(digitLetters2), digitLetters2, "")
	return words
}
