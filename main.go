package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	result := twoSumSolution2(numbers, 9)
	fmt.Println(result)
}

func twoSumSolution1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSumSolution2(nums []int, target int) []int {
	mapHelper := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		diff := target - current
		if j, ok := mapHelper[diff]; ok {
			return []int{j, i}
		}

		mapHelper[current] = i
	}

	return []int{}
}
