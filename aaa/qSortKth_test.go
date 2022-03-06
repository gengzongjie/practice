package aaa

import (
	"fmt"
	"testing"
)

func Test_qSortKth(t *testing.T)  {
	var list = []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	var Kth int
	qSortKth(&list, 0, len(list)-1, 3, &Kth)
	t.Log(Kth)
}

func qSortKth(arr *[]int, begin, end int, kTh int, value *int)  {
	if begin >= end {return}

	pIdx := partitionD(arr, begin, end)
	if pIdx == kTh - 1 {
		*value = (*arr)[pIdx]
		return
	}
	fmt.Scan()
	qSortKth(arr, begin, pIdx-1, kTh, value)
	qSortKth(arr, pIdx+1, end, kTh, value)

	return
}

func partitionD(arr *[]int, begin, end int) int {
	var pIdx = end
	for begin != end {
		for pIdx != begin {
			if (*arr)[begin] >= (*arr)[pIdx] {
				begin ++
			} else {
				(*arr)[begin], (*arr)[pIdx] = (*arr)[pIdx], (*arr)[begin]
				pIdx = begin
			}
		}
		for pIdx != end {
			if (*arr)[pIdx] >= (*arr)[end] {
				end --
			} else {
				(*arr)[pIdx], (*arr)[end] = (*arr)[end], (*arr)[pIdx]
				pIdx = end
			}
		}
	}

	return pIdx
}
