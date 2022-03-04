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

type charCntList struct {
	List []charCnt
}
func (cl charCntList) Len() int {
	return len(cl.List)
}
func (cl charCntList) Less(i, j int) bool {
	return cl.List[i].Cnt > cl.List[j].Cnt
}
func (cl charCntList) Swap(i, j int) {
	cl.List[i], cl.List[j] = cl.List[j], cl.List[i]
}

func charSts(input string) string {
	if len(input) == 0 {return ""}

	var charCntMap = make(map[string]charCnt)

	for _, c := range input {
		char := string(c)
		if v, ok := charCntMap[char]; ok {
			v.Cnt ++
			charCntMap[char] = v
		} else {
			charCntMap[char] = charCnt{
				Char: char,
				Cnt: 1,
			}
		}
	}

	var stsList charCntList
	for _, v := range charCntMap {
		charCntNode := v
		stsList.List = append(stsList.List, charCntNode)
	}

	sort.Sort(stsList)

	var output string
	for _, l := range stsList.List {
		output += fmt.Sprintf("%s%d", l.Char, l.Cnt)
	}

	return output
}
