package problem0001

type candidate struct {
	index int
	value int
}

func twoSum(nums []int, target int) []int {
	var result []int
	for i, v := range nums {
		var cand = candidate{index: i, value: v}
		for j := cand.index + 1; j < len(nums); j++ {
			if cand.value+nums[j] == target {
				result = []int{cand.index, j}
			}
		}
	}
	return result
}

/*
func twoSum2(nums []int, target int) []int {
	var refNums map[int]int
	for i, v := range nums {
		var candVal = target - v
		if j, ok := refNums[candVal]; ok {
			return []int{j, i}
		}
		refNums[v] = i
	}
	return nil
}
*/
