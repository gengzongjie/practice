package aaa

import "testing"

func Test_Qsort(t *testing.T)  {
	var list = []charCnt2{{
		Char: "a",
		Cnt: 1,
	}, {
			Char: "b",
			Cnt: 5,
	}, {
				Char: "c",
				Cnt: 3,
	},
	{
					Char: "d",
					Cnt: 7,
	},{
						Char: "e",
						Cnt: 4,
		}}

	t.Log(quickSort(list, 0, len(list)-1))
}

type charCnt2 struct {
	Char string
	Cnt int
}

func quickSort(arr []charCnt2, beginIdx int, endIdx int) []charCnt2 {
	var pIdx int
	if beginIdx < endIdx {
		pIdx, arr = partitionDesc(arr, beginIdx, endIdx)
		quickSort(arr, beginIdx, pIdx-1)
		quickSort(arr, pIdx+1, endIdx)
	}

	return arr
}

func partitionDesc(arr []charCnt2, leftIdx int, rightIdx int) (int, []charCnt2) {
	var pIdx = rightIdx
	for leftIdx < rightIdx {
		for leftIdx != pIdx {
			if arr[leftIdx].Cnt >= arr[pIdx].Cnt {
				leftIdx ++
			} else {
				arr[leftIdx], arr[pIdx] = arr[pIdx], arr[leftIdx]
				pIdx = leftIdx
			}
		}
		for rightIdx != pIdx {
			if arr[rightIdx].Cnt <= arr[pIdx].Cnt {
				rightIdx --
			} else {
				arr[rightIdx], arr[pIdx] = arr[pIdx], arr[rightIdx]
				pIdx = rightIdx
			}
		}
	}

	return pIdx, arr
}

func Test_qSort(t *testing.T)  {
	var arr = []int{1, 3, 5, 4, 7, 2, 8}
	qSort(&arr, 0, 6)
	t.Log(arr)
}

func qSort(arr *[]int, begin int, end int) {
	if begin >= end {return}

	pIdx := partitionDesc1(arr, begin, end)

	qSort(arr, begin, pIdx-1)
	qSort(arr, pIdx+1, end)

	return
}

func partitionDesc1(arr *[]int, begin, end int) int {
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
