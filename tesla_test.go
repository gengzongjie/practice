package ext

import "testing"

func Test_a(t *testing.T)  {
	arr := []int{7, 3, 7, 3, 1, 3, 4, 1}
	t.Log(Solution(arr))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) == 0 {
		return 0
	}
	
	pList := pureList(A)
	minLen := len(pList)
	var res = 0
	for i :=0; i < len(A)-minLen; i ++ {
		for gap := minLen; gap < len(A)-i; gap ++ {
			j := i + gap
			checkList := A[i:j]
			if isRight(checkList, pList) {
				if res == 0 || gap < res {
					res = gap
				}
			}
		}
	}
	
	return res
}

func isRight(checkList []int, pList []int) bool {
	var checkMap = make(map[int]struct{})
	for _, n := range checkList {
		if _, ok := checkMap[n]; ok {
			continue
		} else {
			checkMap[n] = struct{}{}
		}
	}

	return len(checkMap) == len(pList)
}

func pureList(a []int) []int {
	var pureMap = make(map[int]struct{})
	for _, n := range a {
		if _, ok := pureMap[n]; ok {
			continue
		} else {
			pureMap[n] = struct{}{}
		}
	}

	var res []int
	for k, _ := range pureMap {
		res = append(res, k)
	}

	return res
}