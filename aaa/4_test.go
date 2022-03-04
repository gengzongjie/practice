package aaa

import "testing"

func Test_4(t *testing.T) {
	graph := map[string][]string {
		"A": {"B"},
		"B": {"C", "D"},
		"C": {"D"},
		"D": {"F"},
		"E": {"D"},
		"F": nil,
	}

	t.Log(compilePath(graph, "A"))
}


func compilePath(graph map[string][]string, node string) []string {
	if len(graph) == 0 {return nil}

	var nodeQ = []string{node}
	var checkedNodeMap = make(map[string]bool)
	var pathDesc []string

	for len(nodeQ) > 0 {
		node := nodeQ[0]
		nodeQ = nodeQ[1:]

		if _, ok := checkedNodeMap[node]; ok {
			continue
		} else {
			checkedNodeMap[node] = true
			pathDesc = append(pathDesc, node)
		}

		subNodes := graph[node]
		if len(subNodes) > 0 {
			for _, node := range subNodes {
				nodeQ = append(nodeQ, node)
			}
		}
	}

	var path []string
	for i := len(pathDesc) - 1; i >= 0; i -- {
		path = append(path, pathDesc[i])
	}

	return path
}
