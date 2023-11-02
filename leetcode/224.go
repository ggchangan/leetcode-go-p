package leetcode

import (
	"strconv"
	"strings"
)

var Calculate = calculate

func calculate(s string) int {
	s = s + "#"
	pri := map[byte]map[byte]int{
		'+': {
			'+': 1,
			'-': 1,
			'(': -1,
			')': 1,
			'#': 1,
		},
		'-': {
			'+': 1,
			'-': 1,
			'(': -1,
			')': 1,
			'#': 1,
		},
		'(': {
			'+': -1,
			'-': -1,
			'(': -1,
			')': 0,
			'#': 1,
		},
		')': {
			'+': 1,
			'-': 1,
			'(': 2,
			')': 1,
			'#': 1,
		},
		'#': {
			'+': -1,
			'-': -1,
			'(': -1,
			')': -1,
			'#': 0,
		},
	}
	numberStack := make([]int, 0)
	opStack := make([]byte, 0)

	for i := 0; i < len(s); {
		c := s[i]
		if c == ' ' {
			i++
			continue
		}
		//操作符
		if _, ok := pri[c]; ok {
			l := len(opStack)
			if l == 0 {
				opStack = append(opStack, c)
				i++
				continue
			}
			topOp := opStack[l-1]
			priLevel := pri[topOp][c]
			numberL := len(numberStack)
			switch priLevel {
			case 1:
				number2 := numberStack[numberL-1]
				number1 := numberStack[numberL-2]
				numberStack = numberStack[:numberL-2]
				op := opStack[l-1]
				opStack = opStack[:l-1]
				result := 0
				if op == '+' {
					result = number1 + number2
				} else if op == '-' {
					result = number1 - number2
				}
				numberStack = append(numberStack, result)
			case -1:
				opStack = append(opStack, c)
				i++
			case 0:
				opStack = opStack[:l-1]
				i++
			}
		} else { //操作数
			var sb strings.Builder
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				sb.WriteByte(s[i])
				i++
			}
			number, _ := strconv.Atoi(sb.String())
			numberStack = append(numberStack, number)
		}
	}

	return numberStack[0]
}
