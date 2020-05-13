package main

import "fmt"

func main() {
	dp()
}

func dp() {
	nums := []int{32, 35, 37, 38, 39, 32}
	num := 6

	result := map[int]bool{}
	result[0] = true
	var maxNum int
	for i := 0; i < len(nums); i++ {
		maxNum = max(maxNum, nums[i])
	}
	for i := 1; i <= num; i++ {
		for j := 0; j < len(nums); j++ {
			for n := 0; n <= maxNum*num; n++ {
				if _, ok := result[n]; ok {
					if n+i*nums[j] <= maxNum*num {
						result[n+i*nums[j]] = true
					}
				}
			}
		}
	}
	fmt.Println(len(result) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
