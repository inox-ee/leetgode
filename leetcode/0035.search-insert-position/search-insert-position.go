package problem0035

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	if nums[right] < target {
		return right + 1
	}

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if left == mid {
				break
			}
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
