package main 

import "fmt"

func searchInsert(nums []int, target int) int {
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

	return low
}

func main() {
	nums := []int{1, 3, 5 ,6}
	target := 7

	fmt.Println(searchInsert(nums, target))
}