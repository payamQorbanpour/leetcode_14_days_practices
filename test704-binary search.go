package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		median := (low + high) / 2

		if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(nums) || nums[low] != target {
		return -1
	}

	return low
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 9

	fmt.Println(nums)
	fmt.Println(search(nums, target))
}