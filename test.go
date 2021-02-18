package main

import "fmt"

func fizzBuzz(nums []int) []int {
	for i, v := range nums {
		switch {
		case v%15 == 0:
			fmt.Printf("%d: FizzBuzz\n", i)
		case v%3 == 0:
			fmt.Printf("%d: Fizz\n", i)
		case v%5 == 0:
			fmt.Printf("%d: Buzz\n", i)
		default:
			fmt.Printf("%d: %d\n", i, v)
		}
	}
	return []int{1, 2, 3}
}

func main() {
	var nums = make([]int, 100)
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + 1
		}
	}
	fizzBuzz(nums)
}
