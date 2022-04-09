package main

import "fmt"

func twoSum(numbers []int, target int) []int {
    var output []int
    for i, _ := range(numbers) {
		for j, _ := range(numbers) {
			if i == j {
				continue
			}
			res := numbers[i] + numbers[j]
			if res == target {
				output = append(output, i+1, j+1)
				return output
			}
		}
	}
    return output
}

func main() {
	numbers := []int{2, 3, 4}
	target := 6
	
	fmt.Println(twoSum(numbers, target))
}