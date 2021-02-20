package problem0003

import "strings"

func contain(slc []string, s string) bool {
	for _, r := range slc {
		if r == s {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	var strs = strings.Split(s, "")
	var max = 0
	if len(strs) == 1 {
		return 1
	}
	for i := range strs {
		for j, cand := range strs[i+1:] {
			if i+j+1 == len(strs) || contain(strs[i:i+j+1], cand) {
				if j+1 > max {
					max = j + 1
				}
				break
			}
		}
	}
	return max
}

// func lengthOfLongestSubstring_(s string) int {
// 	var mp = make(map[string]int)  // ASCII 文字コードに対応して [256]int でも良いが、存在判定がやや面倒なのでこれでよさそう
// 	str := strings.Split(s, "")

// 	maxLen, left, right := 0, 0, 0
// 	for right < len(s) {
// 		if v, ok := mp[str[right]]; ok {
// 			left = max_(v+1, left)
// 		}
// 		maxLen = max_(maxLen, right-left+1)
// 		mp[str[right]] = right
// 		right++
// 	}
// 	return maxLen
// }

// func max_(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
