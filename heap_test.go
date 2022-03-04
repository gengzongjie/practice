package ext

import "testing"

func heapifySmallNoMove(arr []int, count int) []int {
	if len(arr) == 0 {
		return nil
	}

	for heapifyIdx := count/2 - 1; heapifyIdx >= 0; heapifyIdx-- {
		var parentIdx = heapifyIdx
		for {
			minIdx := parentIdx
			leftChildIdx := parentIdx*2 + 1
			rightChildIdx := parentIdx*2 + 2

			if leftChildIdx < count && arr[leftChildIdx] < arr[minIdx] {
				minIdx = leftChildIdx
			}
			if rightChildIdx < count && arr[rightChildIdx] < arr[minIdx] {
				minIdx = rightChildIdx
			}

			if minIdx == parentIdx {
				break
			}
			arr[parentIdx], arr[minIdx] = arr[minIdx], arr[parentIdx]
			parentIdx = minIdx
		}
	}

	return arr
}

func heapifyBigNoMove(arr []int, count int) []int {
	if len(arr) == 0 {
		return nil
	}

	for heapifyIdx := count/2 - 1; heapifyIdx >= 0; heapifyIdx-- {
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

			if parentIdx == maxIdx {
				break
			}
			arr[parentIdx], arr[maxIdx] = arr[maxIdx], arr[parentIdx]
			parentIdx = maxIdx
		}
	}

	return arr
}

func topK(arr []int, k int) []int {
	if len(arr) == 0 {
		return nil
	}

	arr = heapifySmallNoMove(arr, k)

	for checkIdx := k; checkIdx < len(arr); checkIdx++ {
		if arr[checkIdx] < arr[0] {
			continue
		}

		arr[0] = arr[checkIdx]
		heapifySmallNoMove(arr, 3)
	}

	return arr[:k]
}

func Test_topK(t *testing.T) {
	var arr = []int{1, 8, 7, 6, 5, 9, 4, 3}
	t.Log(topK(arr, 3))
}

type Node struct {
	Str string
	Cnt int
}

func numStats(arr []string) map[string]int {
	var numCnt = make(map[string]int)

	for _, data := range arr {
		_, ok := numCnt[data]
		if !ok {
			numCnt[data] = 1
		} else {
			numCnt[data]++
		}
	}

	return numCnt
}

func Test_words(t *testing.T) {
	var words = []string{"a", "b", "c", "d", "e", "f", "g", "e", "c", "a"}
	cntMap := numStats(words)

	var cntNodes []Node
	for k, v := range cntMap {
		cntNodes = append(cntNodes, Node{
			Str: k,
			Cnt: v,
		})
	}

	topK := NodeTopK(cntNodes, 3)

	t.Log(topK)

	return
}

func NodeTopK(nodes []Node, k int) []Node {
	if len(nodes) == 0 {
		return nil
	}

	nodes = heapifyNodeSmall(nodes, k)

	for checkIdx := k; checkIdx < len(nodes); checkIdx++ {
		if nodes[checkIdx].Cnt > nodes[0].Cnt {
			nodes[0], nodes[checkIdx] = nodes[checkIdx], nodes[0]
			nodes = heapifyNodeSmall(nodes, k)
		}
	}

	return nodes[:k]
}

func heapifyNodeSmall(nodes []Node, count int) []Node {
	if len(nodes) == 0 {
		return nil
	}

	for heapifyIdx := count/2 - 1; heapifyIdx >= 0; heapifyIdx-- {
		var parentIdx = heapifyIdx
		for {
			minIdx := parentIdx
			leftChildIdx := parentIdx*2 + 1
			rightChildIdx := parentIdx*2 + 2

			if leftChildIdx < count && nodes[leftChildIdx].Cnt < nodes[minIdx].Cnt {
				minIdx = leftChildIdx
			}
			if rightChildIdx < count && nodes[rightChildIdx].Cnt < nodes[minIdx].Cnt {
				minIdx = rightChildIdx
			}

			if parentIdx == minIdx {
				break
			}
			nodes[parentIdx], nodes[minIdx] = nodes[minIdx], nodes[parentIdx]
			parentIdx = minIdx
		}
	}

	return nodes
}
