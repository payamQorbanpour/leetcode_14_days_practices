package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	for i, _ := range(nums) {
		nums[i] = nums[i] * nums[i]
	}

	sort.Ints(nums)

	return nums
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}

	fmt.Println(sortedSquares(nums))
}