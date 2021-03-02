package problem0033

func search(nums []int, target int) int {
	/*
		// O(n)
		for i, v := range nums {
			if v == target {
				return i
			}
		}
		return -1
	*/
	if len(nums) == 0 {
		return -1
	}
	rotated := indexOfMin(nums)
	nums = append(nums[rotated:], nums[:rotated]...)
	tmpIndex := indexOfTarget(nums, target)
	if tmpIndex < 0 {
		return tmpIndex
	}
	tmpIndex = tmpIndex + rotated
	if tmpIndex >= len(nums) {
		return tmpIndex - len(nums)
	}
	return tmpIndex
}

func indexOfTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var center int
	for {
		if target == nums[left] {
			return left
		}
		if target == nums[right] {
			return right
		}
		center = (left + right) / 2
		if left == center {
			break
		}
		if nums[center] < target {
			left = center
		} else {
			right = center
		}
	}
	return -1
}

func indexOfMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] <= nums[right] {
		return left
	}
	center := (left + right) / 2
	for left != center {
		if nums[left] > nums[center] {
			right = center
		} else {
			left = center
		}
		center = (left + right) / 2
	}
	return right
}
