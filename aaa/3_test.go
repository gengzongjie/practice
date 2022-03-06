package aaa

import (
	"testing"
)

func Test_3(t *testing.T)  {
	var input = []string{"a", "b", "c", "d", "a", "d", "a"}
	t.Log(secondMore(input))
}

type charCnt1 struct {
	Char string
	Cnt int
}

func secondMore(input []string) string {
	var charMap = make(map[string]int)
	for _, c := range input {
		if v, ok := charMap[c]; ok {
			v ++
			charMap[c] = v
		} else {
			charMap[c] = 1
		}
	}

	var itmes = make([]charCnt1, len(charMap))
	for k, v := range charMap {
		itmes = append(itmes, charCnt1{
			Char: k,
			Cnt:  v,
		})
	}
	heap := heapSort(itmes, 2)
	return heap[0].Char

	//var topTow = make([]charCnt1, 2)
	//for k, v := range charMap {
	//	if v > topTow[0].Cnt {
	//		topTow[1] = topTow[0]
	//		topTow[0] = charCnt1{
	//			Char: k,
	//			Cnt:  v,
	//		}
	//	} else if v > topTow[1].Cnt {
	//		topTow[1] = charCnt1{
	//			Char: k,
	//			Cnt:  v,
	//		}
	//	}
	//}
	//
	//return topTow[1].Char
}

func heapSort(arr []int, topK int) []int {
	var smallHeap = make([]int, topK)
	for i := 0; i < topK; i ++ {
		smallHeap[i] = arr[i]
	}
	smallHeap = heapify(smallHeap)

	for i := topK; i < len(arr); i ++ {
		if arr[i] > smallHeap[0] {
			smallHeap[0] = arr[i]
			smallHeap = heapify(smallHeap)
		}
	}

	return smallHeap
}

func heapify(heap []int) []int {
	count := len(heap)
	for i := (count-1)/2; i >= 0; i -- {
		rootIdx := i
		for {
			minIdx := rootIdx
			leftIdx := rootIdx*2+1
			rightIdx := rootIdx*2+2
			if leftIdx < count && heap[leftIdx] < heap[minIdx] {
				minIdx = leftIdx
			}
			if rightIdx < count && heap[rightIdx] < heap[minIdx] {
				minIdx = rightIdx
			}
			if minIdx == rootIdx {break}
			heap[minIdx], heap[rootIdx] = heap[rootIdx], heap[minIdx]
			rootIdx = minIdx
		}
	}
}

////func heapSort(items []charCnt1, topK int) []charCnt1 {
////	if len(items) == 0 {return nil}
////
////	var heap = make([]charCnt1, topK)
////	for i := 0; i < topK; i ++ {
////		heap[i] = items[i]
////	}
////	heap = heapify(heap)
////
////	for j := topK; j < len(items); j ++ {
////		if items[j].Cnt > heap[0].Cnt {
////			heap[0] = items[j]
////			heap = heapify(heap)
////		}
////	}
////
////	return heap
////}
////
////func heapify(items []charCnt1) []charCnt1 {
//	for heapifyIdx := len(items)/2-1; heapifyIdx >= 0; heapifyIdx -- {
//		rootIdx := heapifyIdx
//		for {
//			minIdx := rootIdx
//			leftIdx := rootIdx*2 + 1
//			rightIdx := rootIdx*2 + 2
//			if leftIdx < len(items) && items[leftIdx].Cnt < items[minIdx].Cnt {
//				minIdx = leftIdx
//			}
//			if rightIdx < len(items) && items[rightIdx].Cnt < items[minIdx].Cnt {
//				minIdx = rightIdx
//			}
//			if minIdx == rootIdx {
//				break
//			}
//			items[minIdx], items[rootIdx] = items[rootIdx], items[minIdx]
//			rootIdx = minIdx
//		}
//	}
//
//	return items
//}