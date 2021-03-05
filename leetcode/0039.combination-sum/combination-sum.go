package problem0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	sort.IntSlice(candidates).Sort() // sort.Ints(candidates) で等価の操作が可能
	recur(candidates, []int{}, target, &ans)
	return ans
}

func recur(candidates, tmp []int, left int, ans *[][]int) {
	for i, v := range candidates {
		if v == left {
			*ans = append(*ans, append(tmp, v))
			return
		} else if v < left {
			if left >= 2*v {
				// ! tmp = tmp[:len(tmp):len(tmp)] // この１行を入れたら通る！！！
				recur(candidates[i:], append(tmp, v), left-v, ans)
			}
		} else {
			break
		}
	}
	return
}

func sumIntSlice(s []int) int {
	result := 0
	for _, v := range s {
		result += v
	}
	return result
}
