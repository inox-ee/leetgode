package problem0006

import (
	"sort"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var row = 1
	var pn = -1
	var result = ""
	var mp = make(map[int][]int)
	str := strings.Split(s, "")
	for i := range str {
		mp[row] = append(mp[row], i)
		stepRow(&row, &pn, numRows)
	}
	for _, v := range sortMap(mp) {
		for _, vv := range mp[v] {
			result += str[vv]
		}
	}
	return result
}

func sortMap(mp map[int][]int) []int {
	slice1 := make([]int, len(mp))
	index := 0
	for key := range mp {
		slice1[index] = key
		index++
	}
	sort.Ints(slice1)
	return slice1
}

func stepRow(r *int, pn *int, numRows int) {
	if *r >= numRows || *r <= 1 {
		*pn *= -1
	}
	*r += *pn
}
