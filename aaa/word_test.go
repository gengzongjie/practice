package aaa

import "testing"

func Test_Word(t *testing.T)  {
	wordList := []string{"hot","dot","dog","lot","log", "cog"}
	beginWord := "hit"
	endWord := "cog"

	res := ladderLength(beginWord, endWord, wordList)
	t.Log(res)
}

func ladderLength(beginWord string, endWord string, wordList []string) []string {
	if len(beginWord) != len(endWord) {return nil}
	if len(beginWord) == 0 || len(endWord) == 0 {return nil}

	var wordPath []string
	wordPath = append(wordPath, beginWord)
findNext:
	for _, word := range wordList {
		if wordInList(word, wordPath) {continue}
		if wordMatch(word, wordPath[len(wordPath) - 1]) {
			wordPath = append(wordPath, word)
			if wordMatch(endWord, word) {
				wordPath = append(wordPath, endWord)
				return wordPath
			}
			goto findNext
		}
	}
	return nil
}

func wordInList(word string, wordList []string) bool {
	for _, w := range wordList {
		if word == w {
			return true
		}
	}

	return false
}

func Test_wordInList(t *testing.T)  {
	word := "ldd"
	wordList := []string{"hot","dot","dog","lot","log","cog"}

	t.Log(wordInList(word, wordList))
}

func wordMatch(word1 string, word2 string) bool {
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

func Test_wordMatch(t *testing.T)  {
	word1 := "hot"
	word2 := "hog"

	t.Log(wordMatch(word1, word2))
}