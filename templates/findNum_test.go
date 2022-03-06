package templates

import "testing"

func Test_arr(t *testing.T) {
	arr := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	//arr := [][]int{}
	num := 3
	t.Log(FindNum(arr, num))
}

func FindNum(arr [][]int, num int) bool {
	if arr == nil {
		return false
	}
	if len(arr) == 0 {
		return false
	}
	rows := len(arr)
	columns := len(arr[0])
	for row, column := rows-1, 0; row >= 0 && column < columns; {
		compareNum := arr[row][column]
		if compareNum == num {
			return true
		}
		if compareNum > num {
			row--
		} else {
			column++
		}
	}

	return false
}
