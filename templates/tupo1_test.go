package templates

import "testing"

func Test_tupo1(t *testing.T)  {
	var dependencies = map[string][]string{
		"A": {"B"},
		"B": {"C", "D"},
		"C": {"D"},
		"D": {"F"},
		"E": {"D"},
		"F": {},
	}
	t.Log(tupo1(dependencies))
}

func tupo(depends map[string][]string) bool {
	var graph = make(map[string][]string)
	var inDegrees = make(map[string]int)
	var queue []string
	var visited = make(map[string]bool)

	for k, v := range depends {
		inDgree := len(v)
		if inDgree == 0 {
			queue = append(queue, k)
		} else {
			inDegrees[k] = inDgree
		}

		for _, node := range v {
			if n, ok := graph[node]; ok {
				n = append(n, k)
				graph[node] = n
			} else {
				graph[node] = []string{k}
			}
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = true

		for _, subNode := range graph[node] {
			inDegrees[subNode] --

			if inDegrees[subNode] == 0 {
				queue = append(queue, subNode)
			}
		}

	}

	if len(visited) == len(depends) {
		return false
	}

	return true
}

func tupo1(depends map[string][]string) bool {
	var inDegreeTable = make(map[string]int)
	var pathGraph = make(map[string][]string)
	var visited = make(map[string]bool)
	var zeroInDegreeQ []string

	for k, v := range depends {
		inDegree := len(v)
		if inDegree == 0 {
			zeroInDegreeQ = append(zeroInDegreeQ, k)
		} else {
			inDegreeTable[k] = inDegree
		}

		for _, node := range v {
			if graphNode, ok := pathGraph[node]; ok {
				graphNode = append(graphNode, k)
				pathGraph[node] = graphNode
			} else {
				pathGraph[node] = []string{k}
			}
		}
	}

	for len(zeroInDegreeQ) > 0 {
		zeroInDegreeNode := zeroInDegreeQ[0]
		zeroInDegreeQ = zeroInDegreeQ[1:]

		if _, ok := visited[zeroInDegreeNode]; ok {
			continue
		} else {
			visited[zeroInDegreeNode] = true
		}
		if subNodes, ok := pathGraph[zeroInDegreeNode]; ok {
			for _, subNode := range subNodes {
				inDegreeTable[subNode] --

				if inDegreeTable[subNode] == 0 {
					zeroInDegreeQ = append(zeroInDegreeQ, subNode)
				}
			}
		}
	}

	if len(visited) == len(depends) {
		return false
	}

	return true
}