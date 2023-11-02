package leetcode

import "strings"

var Convert = convert

func convert(s string, numRows int) (ans string) {
	if numRows == 1 { //边界条件
		ans = s
		return
	}
	strs := make([]strings.Builder, numRows)

	row := 0
	flag := -1
	for i := 0; i < len(s); i++ {
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		strs[row].WriteByte(s[i])
		row += flag
	}

	for i := 0; i < numRows; i++ {
		ans += strs[i].String()
	}
	return
}
