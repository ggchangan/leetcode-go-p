package leetcode

var RomanToInt = romanToInt

func romanToInt2(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sp := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	if len(s) == 1 {
		return m[s[0]]
	}

	sum := 0
	var i int
	for i = 1; i < len(s); i++ {
		if m[s[i]] <= m[s[i-1]] {
			sum += m[s[i-1]]
		} else {
			sum += sp[s[i-1:i+1]]
			i++
		}
	}

	if i == len(s) {
		sum += m[s[i-1]]
	}

	return sum
}

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}
	}

	return sum
}
