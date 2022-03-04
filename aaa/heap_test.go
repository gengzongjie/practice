package aaa

import (
	"fmt"
	"testing"
)

func createBigHeap(ints []int) []int {
	if ints == nil {return nil}

	var heap = make([]int, len(ints) + 1)

	var count = 0
	for _, v := range ints {
		count ++
		heap[count] = v

		// check if need adjust
		var i = count
		for i/2 > 0 && heap[i] > heap[i/2] {
			heap[i], heap[i/2] = heap[i/2], heap[i]
			i = i/2
		}

	}

	return heap
}

func Test_createBigHeap(t *testing.T)  {
	var ints = []int{50,10,90,30,70,40,80,60,20}
	ints = []int{7,5,19,8,4,1,20,13,16}
	//t.Log(heapifySmallNoMove(ints, len(ints))) // [10 30 40 20 70 90 80 60 50]
	t.Log(heapifyBigNoMove(ints, len(ints)))
	//t.Log(createBigHeap(ints))
	//t.Log(createSmallHeap(ints))
	//t.Log(heapifyBigNoMove(ints))
	//t.Log(heapifyBigNoMove2(ints, len(ints)))
	//t.Log(heapRank(ints))
	//t.Log(TopK(ints, 3))
}

func createSmallHeap(ints []int) []int  {
	if len(ints) == 0 {return nil}

	var heap = make([]int, len(ints) + 1)
	var count = 0

	for _, data := range ints {
		count ++
		heap[count] = data

		// check and adjust
		var i = count
		for i/2 >0 && heap[i] < heap[i/2] {
			heap[i], heap[i/2] = heap[i/2], heap[i]
			i = i/2
		}
	}

	return heap
}

/*
自上而下，原地堆化
从最后一个非叶子节点开始向上堆化，每次堆化都要把变动的子树从上向下全部堆化
 */

func heapifyBigNoMove(arr []int, count int) []int {
	if len(arr) == 0 {
		return nil
	}

	for heapifyIdx := count/2; heapifyIdx >= 0; heapifyIdx -- {
		var parentIdx = heapifyIdx
		for {
			maxIdx := parentIdx
			leftChildIdx := parentIdx*2 + 1
			rightChildIdx := parentIdx*2 + 2

			if leftChildIdx < count && arr[leftChildIdx] > arr[maxIdx] {
				maxIdx = leftChildIdx
			}
			if rightChildIdx < count && arr[rightChildIdx] > arr[maxIdx] {
				maxIdx = rightChildIdx
			}

			if maxIdx == parentIdx {
				break
			}
			arr[parentIdx], arr[maxIdx] = arr[maxIdx], arr[parentIdx]
			parentIdx = maxIdx
		}
	}

	return arr
}

func heapRank(arr []int) []int {
	var count = len(arr)
	arr = heapifyBigNoMove(arr, count)
	for count > 1 {
		arr[0], arr[count-1] = arr[count-1], arr[0]
		count -= 1
		heapifyBigNoMove(arr, count)
	}

	return arr
}

func heapifySmallNoMove(arr []int, count int) []int {
	if len(arr) == 0 {
		return nil
	}

	for heapifyIdx := count/2; heapifyIdx >= 0; heapifyIdx -- {
		var parentIdx = heapifyIdx
		for {
			minIdx := parentIdx
			leftChildIdx := parentIdx*2+1
			rightChildIdx := parentIdx*2+2

			if leftChildIdx < count && arr[leftChildIdx] < arr[minIdx] {
				minIdx = leftChildIdx
			}
			if rightChildIdx < count && arr[rightChildIdx] < arr[minIdx] {
				minIdx = rightChildIdx
			}
			if minIdx == parentIdx {break}
			arr[parentIdx], arr[minIdx] = arr[minIdx], arr[parentIdx]
			parentIdx = minIdx
		}
	}

	return arr
}

func TopK(arr []int, k int) []int  {
	if len(arr) == 0 {return nil}

	arr = heapifySmallNoMove(arr, k)

	for i := k; i < len(arr); i ++ {
		if arr[i] <= arr[0] {continue}

		// swap and re heapify
		arr[0], arr[i] = arr[i], arr[0]
		arr = heapifySmallNoMove(arr, k)
	}

	fmt.Println(arr)
	return arr[:k]
}