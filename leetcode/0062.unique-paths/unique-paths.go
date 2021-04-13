package problem0062

var memo = make(map[[2]int]int)

func uniquePaths(m int, n int) int {

	return dfs(m-1, n-1)
}

func dfs(up, down int) int {
	res := 0
	if up == 0 || down == 0 {
		return res + 1
	}
	if v, ok := memo[[2]int{up, down}]; ok {
		return v
	}
	res += dfs(up-1, down)
	res += dfs(up, down-1)
	memo[[2]int{up, down}] = res
	return res
}
