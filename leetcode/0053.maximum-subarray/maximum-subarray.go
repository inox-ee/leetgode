package problem0053

import (
	"math"
)

func maxSubArray(nums []int) int {
	max := -1*int(math.Pow10(5)) - 1

	for left := 0; left < len(nums); left++ {
		for right := left; right < len(nums); right++ {
			tmpSum := sum(nums[left : right+1])
			if tmpSum > max {
				max = tmpSum
			}
		}
	}
	return max
}

func sum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	current_subarray := []int{nums[0]}
	max_subarray := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		a := sum(current_subarray) + nums[i]
		if a < nums[i] {
			current_subarray = make([]int, 1, 1)
			current_subarray = []int{nums[i]}
		} else {
			current_subarray = append(current_subarray, nums[i])
		}
		if sum(current_subarray) > sum(max_subarray) {
			max_subarray = current_subarray
		}
	}

	return sum(max_subarray)
}
