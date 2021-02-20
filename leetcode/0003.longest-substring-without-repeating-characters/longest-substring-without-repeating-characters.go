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

func lengthOfLongestSubstring_(s string) int {
	var location [256]int // ASCII文字コードが 256 なので (普通にmapにしても良いが、場合分けが若干面倒。found/not found ではなく、インデックスとの大小比較がしたいので。)

	maxLen, left, right := 0, 0, 0

	for left < len(s) {
		if right+1 < len(s) && location[s[right+1]-'a'] <= 0 {
			location[s[right+1]-'a']++
			right++
		} else {
			location[s[left]-'a']--
			left++
		}
		maxLen = max_(maxLen, right-left+1)
	}
	return maxLen
}

func max_(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
