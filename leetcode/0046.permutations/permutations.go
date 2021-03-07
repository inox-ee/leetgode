package problem0046

func permute(nums []int) [][]int {
	var result [][]int
	dfs(nums, []int{}, &result)
	return result
}

func dfs(left, solution []int, result *[][]int) {
	if len(left) == 0 {
		*result = append(*result, solution)
		return
	}
	for i, v := range left {
		solution = solution[:len(solution):len(solution)]
		left = left[:len(left):len(left)]
		dfs(append(left[:i], left[i+1:]...), append(solution, v), result)
	}
	return
}

/*
func permute2(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	dfs2([]int{}, nums, &ret)
	return ret
}

func dfs2(pat, nums []int, ret *[][]int) {
	if len(nums) == 0 {
		*ret = append(*ret, pat)
		return
	}
	for i := 0; i < len(nums); i++ {
		dfs2(append(pat, nums[i]), exclude(nums, i), ret)
	}
}

func exclude(nums []int, i int) []int { // これは今後も流用した方がいいかも
	ret := make([]int, 0, len(nums)-1)
	ret = append(ret, nums[:i]...)
	ret = append(ret, nums[i+1:]...)
	return ret
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}
*/
