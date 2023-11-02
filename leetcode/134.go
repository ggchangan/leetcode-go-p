package leetcode

var CanCompleteCircuit = canCompleteCircuit

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; {
		sumOfCost, sumOfGas, cnt := 0, 0, 0
		for cnt < l {
			sumOfCost += cost[(i+cnt)%l]
			sumOfGas += gas[(i+cnt)%l]
			if sumOfGas < sumOfCost {
				break
			}
			cnt++
		}

		if cnt == l {
			return i
		} else {
			i = i + cnt + 1
		}
	}

	return -1
}
