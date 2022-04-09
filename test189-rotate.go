package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	if k == 0 || k == len(nums) {
		return
	}
	if k > len(nums) {
		k %= len(nums)
	}
	reverse(nums[0 : len(nums)-k])
	fmt.Println(nums)
	reverse(nums[len(nums)-k:])
	fmt.Println(nums)
	reverse(nums)

	fmt.Println(nums)
}

func reverse(s []int) {
	if len(s) == 0 {
		return
	}
	end := len(s) - 1
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[end-i] = s[end-i], s[i]
	}
}


func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 1
	
	rotate(nums, k)
}