package problem0001

func _twoSum(nums []int, target int) []int {
	var mp = make(map[int]int)
	var result []int

	for index, num := range nums {
		cand := target - num
		if i, ok := mp[cand]; ok {
			result = []int{i, index}
		} else {
			mp[num] = index
		}
	}
	return result
}

// func main() {
// 	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
// 	fmt.Printf("%v\n", twoSum([]int{3, 2, 4}, 6))
// 	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))
// }
