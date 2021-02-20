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

func lengthOfLongestSubstring(s string) int {
	var location [256]int // ASCII文字コードが 256 なので (普通にmapにしても良いが、場合分けが若干面倒。found/not found ではなく、インデックスとの大小比較がしたいので。)

	maxLen, left, right := 0, 0, 0

	for left < len(s) {
		if right+1 < len(s) && location[s[right+1]-'a'] <= 0 {
			location[s[right+1]-'a']++
			right++
		} else {
			location[s[left]-'a']--
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(lengthOfLongestSubstring("abcabacbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring("pwwke"))
	println(lengthOfLongestSubstring("au"))
	println(lengthOfLongestSubstring("bbbbbb"))
	println(lengthOfLongestSubstring(" "))
	println(lengthOfLongestSubstring(""))
}
