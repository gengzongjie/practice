package templates

import (
"testing"
)

func Test_package(t *testing.T) {
	var weights = []int{3, 4, 7, 8, 9}
	var values = []int{4, 5, 10, 11, 13}
	validDp := knapsackDp(weights, values, 17)
	output := outputKnapsackDp(validDp, weights)
	t.Log(validDp)
	t.Log(output)
}

func outputKnapsackDp(valueDp [][]int, weights []int) []int {
	count := len(valueDp)
	maxWeight := len(valueDp[0]) - 1
	var output = make([]int, count)

	var curWeight = maxWeight
	for i := count - 1; i > 0; i-- {
		if valueDp[i][curWeight] > valueDp[i-1][curWeight] {
			output[i] = 1
			curWeight -= weights[i]
		} else {
			output[i] = 0
		}
	}
	if curWeight == 0 {
		output[0] = 0
	} else {
		output[0] = 1
	}

	return output
}

func knapsackDp(weights []int, values []int, maxWeight int) [][]int {
	rowMax := len(weights)
	colMax := maxWeight + 1
	var valueDp = make([][]int, rowMax)
	for i := 0; i < rowMax; i++ {
		valueDp[i] = make([]int, colMax)
	}

	for w := 0; w < colMax; w++ {
		if w < weights[0] {
			valueDp[0][w] = 0
		} else {
			valueDp[0][w] = values[0]
		}
	}

	for i := 1; i < rowMax; i++ {
		valueDp[i][0] = 0
		weight := weights[i]
		value := values[i]
		for w := 1; w < colMax; w++ {
			curOutMax := valueDp[i-1][w]
			if w >= weight {
				curInMax := value + valueDp[i-1][w-weight]
				if curInMax > curOutMax {
					valueDp[i][w] = curInMax
					continue
				}
			}
			valueDp[i][w] = curOutMax
		}
	}

	return valueDp
}
