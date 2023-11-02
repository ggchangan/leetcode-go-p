package leetcode

import (
	"fmt"
)

var AddBinary = addBinary

func addBinary(a string, b string) (ans string) {
	l := 0
	al := len(a)
	bl := len(b)
	if al >= bl {
		l = al
	} else {
		l = bl
	}

	carry := 0
	for i := 0; i < l; i++ {
		if i < al {
			//al-i-1
			carry += int(a[al-i-1] - '0')
		}
		if i < bl {
			//bl-i-1
			carry += int(b[bl-i-1] - '0')
		}
		ans = fmt.Sprintf("%d", carry%2) + ans
		carry /= 2
	}
	if carry != 0 {
		ans = "1" + ans
	}

	return
}
