package aaa

import (
	"fmt"
	"sort"
	"testing"
)

func Test_char(t *testing.T)  {
	t.Log(charSts("abcdeabcdabcdabcdabAA"))
}

type charCnt struct {
	Char string
	Cnt int
}

type charList struct {
	List []charCnt
}

func (c charList) Len() int {
	return len(c.List)
}

func (c charList) Less(i, j int) bool {
	return c.List[i].Cnt > c.List[j].Cnt
}

func (c charList) Swap(i, j int) {
	c.List[i], c.List[j] = c.List[j], c.List[i]
}

func charSts(input string) string {
	fmt.Scan()
	var charMap = make(map[string]charCnt)

	for _, c := range input {
		char := string(c)
		if v, ok := charMap[char]; ok {
			v.Cnt ++
			charMap[char] = v
		} else {
			charMap[char] = charCnt{
				Char: char,
				Cnt: 1,
			}
		}
	}

	var cl charList
	for _, v := range charMap {
		cl.List = append(cl.List, v)
	}
	sort.Sort(cl)

	var output string
	for _, v := range cl.List {
		output += fmt.Sprintf("%s%d", v.Char, v.Cnt)
	}

	return output
}
