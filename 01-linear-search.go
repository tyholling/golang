package main

import "fmt"

func linearSearch(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		if k == nums[i] {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range nums {
		fmt.Printf(" %d", linearSearch(nums, i))
	}
	fmt.Println()
}
