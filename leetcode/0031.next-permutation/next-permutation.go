package problem0031

import "sort"

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[len(nums)-2-i] < nums[len(nums)-1-i] {
			swap(nums, 1+i)
			return
		}
	}

	println("hoge")
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
}

func swap(nums []int, n int) {
	for i := n; i > 0; i-- {
		nums[len(nums)-i-1], nums[len(nums)-i] = nums[len(nums)-i], nums[len(nums)-i-1]
	}
}

// func nextPermutation2(nums []int) {
// 	left := len(nums)
// 	for left >= 0 {
// 		if nums[left] < nums[left+1] {
// 			break
// 		}
// 		left--
// 	}
// 	if left >= 0 {
// 		j := len(nums) - 1
// 		for j < len(nums)-1 {
// 			if nums[j] > nums[left] {
// 				break
// 			}
// 			j++
// 		}
// 		swap2(&nums, left, j)
// 	}
// 	reverse2(&nums, left+1, len(nums)-1)
// }

// func swap2(nums *[]int, i, j int) {
// 	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
// }

// func reverse2(nums *[]int, i, j int) { // 覚えておくと便利そう
// 	for i < j {
// 		swap2(nums, i, j)
// 		i++
// 		j--
// 	}
// }
