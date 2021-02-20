package main

import (
	"fmt"
	"reflect"
	"strings"
)

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

func create0To99() []int {
	var nums = make([]int, 100)
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + 1
		}
	}
	return nums
}

// 	readString("abcdef~A%あ👩👳‍♀️😀😊🤩🙄☺")
func readString(input string) {
	for _, c := range input {
		fmt.Printf("%d %c ", c, c)
		fmt.Println(reflect.TypeOf(c))
	}

	fmt.Println("---")

	slice := strings.Split(input, "")
	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func main() {
}
