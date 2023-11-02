package leetcode

var LetterCombinations = letterCombinations

func letterCombinations(digits string) []string {
	dL := len(digits)
	if dL == 0 {
		return []string{}
	}
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
	var words []string
	var combination func(l int, word string)
	combination = func(l int, word string) {
		if l == dL {
			words = append(words, word)
			return
		}

		for _, c := range digitLetters[string(digits[l])] {
			combination(l+1, word+string(c))
		}
	}
	words = make([]string, 0)
	combination(0, "")
	return words
}
