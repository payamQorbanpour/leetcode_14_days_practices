package main

import "fmt"

var bad int = 4

func firstBadVersion(n int) int {
	low := 0
	high := n - 1

	for low <= high {
		median := (low + high) / 2

		if isBadVersion(median) == false {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return low
}

func main() {
	n := 5
	fmt.Println(firstBadVersion(n))
}