package problem0003

import "strings"

func _lengthOfLongestSubstring(s string) int {
	var mp = make(map[string]int)
	str := strings.Split(s, "")

	maxLen, left, right := 0, 0, 0
	for right < len(s) {
		if v, ok := mp[str[right]]; ok {
			left = max(v+1, left)
		}
		maxLen = max(maxLen, right-left+1)
		mp[str[right]] = right
		right++
	}
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
