package leetcode

import "strconv"

var EvalRPN = evalRPN

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	op := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	for i := range tokens {
		if op[tokens[i]] {
			l := len(stack)
			number2 := stack[l-1]
			number1 := stack[l-2]
			stack = stack[:l-2]
			result := 0
			switch tokens[i] {
			case "+":
				result = number1 + number2
			case "-":
				result = number1 - number2
			case "*":
				result = number1 * number2
			case "/":
				result = number1 / number2
			}
			stack = append(stack, result)
		} else {
			number, _ := strconv.Atoi(tokens[i])
			stack = append(stack, number)
		}
	}

	return stack[0]
}
