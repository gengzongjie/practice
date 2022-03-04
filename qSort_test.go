package ext

import (
	"testing"
)

func Test_qSort(t *testing.T) {
	var list = []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	t.Log(qSortDesc(list, 0, len(list)-1))
}

func qSortDesc(arr []int, beginIdx int, endIdx int) []int {
	var pIdx int

	if beginIdx < endIdx {
		pIdx, arr = partitionDesc(arr, beginIdx, endIdx)

		qSortDesc(arr, beginIdx, pIdx-1)
		qSortDesc(arr, pIdx+1, endIdx)
	}

	return arr
}

func partitionDesc(arr []int, beginIdx int, endIdx int) (int, []int) {
	var pIdx = endIdx
	leftIdx := beginIdx
	rightIdx := endIdx
	for leftIdx != rightIdx {
		// check left
		for leftIdx != pIdx {
			if arr[leftIdx] >= arr[pIdx] {
				leftIdx++
			} else {
				arr[leftIdx], arr[pIdx] = arr[pIdx], arr[leftIdx]
				pIdx = leftIdx
			}
		}
		// check right
		for rightIdx != pIdx {
			if arr[rightIdx] <= arr[pIdx] {
				rightIdx--
			} else {
				arr[rightIdx], arr[pIdx] = arr[pIdx], arr[rightIdx]
				pIdx = rightIdx
			}
		}
	}

	return pIdx, arr
}

func Test_qSortDesc2(t *testing.T) {
	var list = []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	//qSortDesc2(&list, 0, len(list) - 1)
	v := qSortDesc2(&list, 0, len(list)-1, 3)

	t.Log(v)
}

func qSortDesc2(arr *[]int, beginIdx int, endIdx int, k int) int {
	if beginIdx >= endIdx {
		return 0
	}

	pIdx := partitionDesc2(arr, beginIdx, endIdx)
	if pIdx == k-1 {
		return (*arr)[pIdx]
	}

	qSortDesc2(arr, beginIdx, pIdx-1, k)
	qSortDesc2(arr, pIdx+1, endIdx, k)

	return 0
}

func partitionDesc2(arr *[]int, leftIdx int, rightIdx int) int {
	pIdx := rightIdx

	for leftIdx < rightIdx {
		for leftIdx != pIdx {
			if (*arr)[leftIdx] >= (*arr)[pIdx] {
				leftIdx++
			} else {
				(*arr)[leftIdx], (*arr)[pIdx] = (*arr)[pIdx], (*arr)[leftIdx]
				pIdx = leftIdx
			}
		}

		for rightIdx != pIdx {
			if (*arr)[rightIdx] <= (*arr)[pIdx] {
				rightIdx--
			} else {
				(*arr)[rightIdx], (*arr)[pIdx] = (*arr)[pIdx], (*arr)[rightIdx]
				pIdx = rightIdx
			}
		}

	}

	return pIdx
}

func Test_qsortAsc(t *testing.T) {
	var list = []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	qSortAsc(&list, 0, len(list)-1)

	t.Log(list)
}

func qSortAsc(arr *[]int, beginIdx int, endIdx int) {
	if beginIdx < endIdx {
		pIdx := partitionAsc(arr, beginIdx, endIdx)

		qSortAsc(arr, beginIdx, pIdx-1)
		qSortAsc(arr, pIdx+1, endIdx)
	}

	return
}

func partitionAsc(arr *[]int, leftIdx int, rightIdx int) int {
	pIdx := rightIdx

	for leftIdx < rightIdx {
		for leftIdx != pIdx {
			if (*arr)[leftIdx] <= (*arr)[pIdx] {
				leftIdx++
			} else {
				(*arr)[leftIdx], (*arr)[pIdx] = (*arr)[pIdx], (*arr)[leftIdx]
				pIdx = leftIdx
			}
		}
		for rightIdx != pIdx {
			if (*arr)[rightIdx] >= (*arr)[pIdx] {
				rightIdx--
			} else {
				(*arr)[rightIdx], (*arr)[pIdx] = (*arr)[pIdx], (*arr)[rightIdx]
				pIdx = rightIdx
			}
		}
	}

	return pIdx
}

func Test_FindKthBig(t *testing.T) {
	var list = []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	var k = 3

	var pIdx int
	var beginIdx = 0
	var endIdx = len(list) - 1
	for beginIdx < endIdx {
		pIdx = partitionDesc2(&list, beginIdx, endIdx)
		if pIdx == k-1 {
			break
		}
		if pIdx > k-1 {
			endIdx = pIdx - 1
		} else {
			beginIdx = pIdx + 1
		}
	}

	t.Log(list[pIdx])
}
