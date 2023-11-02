package leetcode

import "strings"

var IntToRoman = intToRoman

func intToRoman(num int) (s string) {
	m := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}

	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < len(numbers); i++ {
		cnt := num / numbers[i]
		num = num - cnt*numbers[i]
		s += strings.Repeat(m[numbers[i]], cnt)
	}

	return
}
