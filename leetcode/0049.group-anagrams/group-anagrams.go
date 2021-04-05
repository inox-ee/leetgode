/*
- 方法 1
  - 単語をソート
	- 一致するキーがあるか確認、なかったらkeyに追加
	- 一致したら元の値をvalueに追加
*/

package problem0049

import (
	"sort"
)

type sortRunes []rune
type hashMap map[string][]string

// func main() {
// 	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	fmt.Printf("%v\n", groupAnagrams(strs))
// }

func groupAnagrams(strs []string) [][]string {
	m := make(hashMap)
	for _, v := range strs {
		sortedStr := sortStringByCharacter(v)
		if arr, ok := m[sortedStr]; ok {
			m[sortedStr] = append(arr, v)
		} else {
			m[sortedStr] = []string{v}
		}
	}
	return m.Values()
}

func (m hashMap) Values() [][]string {
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortStringByCharacter(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
