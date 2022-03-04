package aaa

import (
	"testing"
)

/*
arrDict = [
    [0, 2, 3, 0],
    [0, 6, 0, 7],
    [3, 5, 0, 3],
    [0, 4, 0, 4],
];
 */

const whiteVal = 0

func Test_2(t *testing.T)  {
	var arrDict = [][]int{
		[]int{0, 2, 3,0},
		[]int{0,6,0,7},
		[]int{3,5,0,3},
		[]int{0,4,0,4},
	}

	t.Log(whiteSts(arrDict))
}

func whiteSts(arr [][]int) []int {
	if len(arr)==0 {
		return nil
	}
	if len(arr[0]) == 0 {
		return nil
	}

	var whiteList [][]int
	rowMax := len(arr)
	colMax := len(arr[0])

	for i := 0; i < rowMax; i ++ {
		for j := 0; j < colMax; j ++ {
			v := arr[i][j]
			if v == 0 {
				whiteList = append(whiteList, []int{i, j})
			}
		}
	}

	var deepMap = make(map[int]bool)
	var checkedNodeMap = make(map[int]bool)
	for _, node := range whiteList {
		rowIdx := node[0]
		colIdx := node[1]
		deep := maxDeep(rowIdx, colIdx, arr, checkedNodeMap)
		deepMap[deep] = true
	}

	var deeps []int
	for k, _ := range deepMap {
		deeps = append(deeps, k)
	}

	return deeps
}

func maxDeep(i, j int, arr [][]int, checkedNodeMap map[int]bool) int {
	subNodes := [][]int{{i-1,j-1},{i-1,j},{i-1,j+1},{i,j-1},{i,j+1},{i+1,j-1},{i+1,j},{i+1,j+1}}
	rowMax := len(arr)
	colMax := len(arr[0])

	checkedNodeMap[i*colMax+j] = true
	for _, node := range subNodes {
		rowIdx := node[0]
		colIdx := node[1]
		if rowIdx >=0 && rowIdx<rowMax && colIdx>=0 && colIdx<colMax && arr[rowIdx][colIdx] == 0 {
			nodeKey := rowIdx*colMax+colIdx
			if _, ok := checkedNodeMap[nodeKey]; ok {
				continue
			}

			return maxDeep(rowIdx, colIdx, arr, checkedNodeMap) + 1
		}
	}

	return 1
}
