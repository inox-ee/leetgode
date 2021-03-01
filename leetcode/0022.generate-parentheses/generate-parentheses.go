package problem0022

func generateParenthesis(n int) []string {
	return []string{}
}

/*
func generateParenthesis(n int) []string {
	var ans []string
	backtrack("", 0, 0, &ans, n)
	return ans
}

func backtrack(s string, left int, right int, result *[]string, n int) {
	if len(s) == 2*n {
		*result = append(*result, s)
		return
	}
	if left < n {
		backtrack(s+"(", left+1, right, result, n)
	}
	if right < left {
		backtrack(s+")", left, right+1, result, n)
	}
}
*/
