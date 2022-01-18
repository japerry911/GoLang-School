package main

import "fmt"

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	ts := 10

	result := TwoNumberSum(a, ts)

	fmt.Println(result)
	// [-1, 11] *order does not matter*
}

// O(n) time & O(n) space
func TwoNumberSum(array []int, target int) []int {
	nums := make(map[int]bool)

	for _, v := range array {
		potentialMatch := target - v

		if _, exists := nums[potentialMatch]; exists {
			return []int{potentialMatch, v}
		}

		nums[v] = true
	}

	return []int{}
}

// Slow way | O(n^2) time & O(1) space
// func TwoNumberSum(array []int, target int) []int {
// 	var total int

// 	for i, v := range array {
// 		total = v
// 		for i2, v2 := range array {
// 			if i == i2 {
// 				continue
// 			} else if total + v2 == target {
// 				return []int{v, v2}
// 			}
// 		}
// 	}

// 	return []int{}
// }
