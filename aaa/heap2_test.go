package aaa

import (
	"testing"
)

type Node struct {
	Value int
	Times int
}

func Test_heap2(t *testing.T)  {
	var arr = []int{-1,1,4,-4,3,5,4,-2,3,-1}
	arr = []int{5,-3,9,1,7,7,9,10,2,2,10,10,3,-1,3,7,-9,-1,3,3}
	t.Log(topKFrequent(t, arr, 3))
}

func topKFrequent(t *testing.T, nums []int, k int) []int {
	if len(nums) == 0 {return nil}

	numTimes := numStats(nums)
	t.Logf("numTimes:%+v", numTimes)

	numTimes = heapifySmallNoMove2(numTimes, k)
	t.Logf("init heap: %+v", numTimes)

	for i := k; i < len(numTimes); i ++ {
		if numTimes[i].Times > numTimes[0].Times {
			numTimes[0], numTimes[i] = numTimes[i], numTimes[0]
			numTimes = heapifySmallNoMove2(numTimes, k)
			t.Logf("%d, heap: %+v",i, numTimes)
		}
	}

	var res []int
	for i := 0; i < k; i ++ {
		res = append(res, numTimes[i].Value)
	}

	return res
}

func numStats(nums []int) []Node {
	var numMap = make(map[int]int)

	for _, num := range nums {
		_, ok := numMap[num]
		if !ok {
			numMap[num] = 1
		} else {
			numMap[num] ++
		}
	}

	var nodes []Node
	for k, v := range numMap {
		node := Node{
			Value: k,
			Times: v,
		}
		nodes = append(nodes, node)
	}

	return nodes
}

func Test_heapifySmallNoMove2(t *testing.T)  {
	var nodes = []Node{
		{
			Value: 10,
			Times: 3,
		},
		{
				Value: -1,
				Times: 2,
		},
		{
					Value: -3,
					Times: 1,
		},
	}

	t.Log(heapifySmallNoMove2(nodes, 3))
}

func heapifySmallNoMove2(nodes []Node, k int) []Node {
	if len(nodes) == 0 {return nil}

	for heapifyIdx := k/2; heapifyIdx >= 0; heapifyIdx -- {
		var parentIdx = heapifyIdx
		for {
			minIdx := parentIdx
			leftChildIdx := parentIdx*2 + 1
			rightChildIdx := parentIdx*2 + 2

			if leftChildIdx < k && nodes[leftChildIdx].Times < nodes[minIdx].Times {
				minIdx = leftChildIdx
			}
			if rightChildIdx < k && nodes[rightChildIdx].Times < nodes[minIdx].Times {
				minIdx = rightChildIdx
			}

			if minIdx == parentIdx {break}
			nodes[parentIdx], nodes[minIdx] = nodes[minIdx], nodes[parentIdx]
			parentIdx = minIdx
		}
	}

	return nodes
}
