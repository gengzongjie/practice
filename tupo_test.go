package ext

import "testing"

func Test_Topo(t *testing.T) {
	var dependencies = map[string][]string{
		"A": {"B"},
		"B": {"C", "D"},
		"C": {"D"},
		"D": {"F"},
		"E": {"D"},
		"F": {},
	}
	t.Log(topo(t, dependencies, 6))
}

func topo(t *testing.T, dependencies map[string][]string, num int) bool {
	if dependencies == nil {
		return false
	}

	var inDegrees = make(map[string]int)
	var graph = make(map[string][]string)
	var queue []string
	var visited []string

	for k, v := range dependencies {
		if len(v) == 0 {
			queue = append(queue, k)
		} else {
			inDegrees[k] = len(v)
		}

		for _, dependency := range v {
			if _, ok := graph[dependency]; ok {
				graph[dependency] = append(graph[dependency], k)
			} else {
				graph[dependency] = []string{k}
			}
		}
	}
	t.Log(inDegrees, graph, queue)

	for len(queue) > 0 {
		zeroDegreeNode := queue[0]
		queue = queue[1:]

		visited = append(visited, zeroDegreeNode)
		if zeroDegreeNode == "B" {
			t.Logf("find: %v", visited)
		}

		for _, node := range graph[zeroDegreeNode] {
			inDegrees[node]--

			if inDegrees[node] == 0 {
				queue = append(queue, node)
			}
		}
	}
	t.Log(visited)

	if len(visited) != num {
		return true
	}

	return false
}
