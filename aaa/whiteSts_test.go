package aaa

import "testing"

func Test_whiteSts(t *testing.T)  {
	var arrDict = [][]int{
		[]int{0, 2, 3,0},
		[]int{0,6,0,7},
		[]int{3,5,0,3},
		[]int{0,4,0,4},
	}

	t.Log(whiteSts2(arrDict))
}

func whiteSts2(input [][]int) []int {
	var whiteNodes [][]int

	var rowMax = len(input)
	if rowMax == 0 {
		return nil
	}
	var colMax = len(input[0])

	for i := 0; i < rowMax; i ++ {
		for j := 0; j < colMax; j ++ {
			if input[i][j] == 0 {
				whiteNodes = append(whiteNodes, []int{i, j})
			}
		}
	}

	var checkNode = make(map[int]bool)
	var deeps = make(map[int]int)
	for _, node := range whiteNodes {
		rowIdx := node[0]
		colIdx := node[1]
		nodeKey := rowIdx*colMax + colIdx
		if _, ok := checkNode[nodeKey]; ok {
			continue
		} else {
			checkNode[nodeKey] = true
			maxDeep := maxDeep2(rowIdx, colIdx, input, checkNode)
			deeps[nodeKey] = maxDeep
		}
	}

	var deepList []int
	for _, v := range deeps {
		deepList = append(deepList, v)
	}

	return deepList
}

func maxDeep2(rowIdx, colIdx int, input [][]int, checkedNode map[int]bool) int {
	var rowMax = len(input)
	var colMax = len(input[0])
	arroundNodes := [][]int{{rowIdx-1, colIdx-1}, {rowIdx-1, colIdx}, {rowIdx-1, colIdx+1}, {rowIdx, colIdx-1}, {rowIdx, colIdx+1}, {rowIdx+1, colIdx-1}, {rowIdx+1, colIdx}, {rowIdx+1, colIdx+1}}
	for _, node := range arroundNodes {
		i := node[0]
		j := node[1]
		if i >=0 && i < rowMax && j >= 0 && j < colMax && input[i][j] == 0 {
			nodeKey := i*colMax + j
			if _, ok := checkedNode[nodeKey]; ok {
				continue
			}
			checkedNode[nodeKey] = true
			return maxDeep2(i, j, input, checkedNode) + 1
		}
	}

	return 1
}
