package templates

import (
	"fmt"
	"testing"
)

func Test_Word2(t *testing.T)  {
	var beginWord, endWord = "hit", "cog"
	var wordList = []string{"hot","dot","dog","lot","log","cog"}
	res := Word2(t, beginWord, endWord, wordList)
	t.Log(res)

	var printStr string
	printStr += "the result is: "
	for i := len(res) - 1; i >= 0; i -- {
		printStr += fmt.Sprintf("%s", res[i])
		if i != 0 {
			printStr += ","
		}
	}
	t.Log(printStr)
}

/*
begin 和 end都要在列表里，建到图里面
 */
func Word2(t *testing.T, beginWord string, endWord string, wordList []string) []string {
	if len(beginWord) != len(endWord) {return nil}
	if beginWord == "" || endWord == "" {return nil}
	graph := buildGraph(append(wordList, beginWord))
	t.Log(graph)

	var res []string
	var queue = []string{beginWord}
	var visited = map[string]struct{}{beginWord: {}}
	var pre = map[string]string{beginWord:""}

	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]
		childs, ok := graph[word]
		if ok {
			for _, child := range childs {
				_, ok := visited[child]
				if ok {continue}
				visited[child] = struct{}{}
				pre[child] = word
				if child == endWord {
					// print path
					for w := child; w != ""; {
						res = append(res, w)
						w = pre[w]
					}
					return res
				}
				queue = append(queue, child)
			}
		}
	}
	
	return nil
}

func buildGraph(wordList []string) map[string][]string  {
	if wordList == nil {return nil}

	var graph = make(map[string][]string)

	for _, word := range wordList {
		var childs []string
		for _, w := range wordList {
			if word == w {continue}
			if wordMatch1(word, w) {
				childs = append(childs, w)
			}
		}
		graph[word] = childs
	}

	return graph
}

func Test_buildGraph(t *testing.T)  {
	var wordList = []string{"hit", "hot","dot","dog","lot","log", "cog"}
	t.Log(buildGraph(wordList))
}

func wordMatch1(word1 string, word2 string) bool {
	if len(word1) != len(word2) {return false}
	var diffCharCnt int
	for i := 0; i < len(word1); i ++ {
		if word1[i] != word2[i] {
			if diffCharCnt == 1 {
				return false
			}
			diffCharCnt ++
		}
	}

	return true
}