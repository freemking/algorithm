package main

import (
	"fmt"
)

func main() {
	dp()
}

func dp() {
	nums := []int{3200, 2800, 2900, 3000, 2700, 2800}
	num := 6

	numsLen := len(nums)
	result := map[int][]int{}
	result[0] = make([]int, numsLen+1)
	var maxNum int
	for i := 0; i < numsLen; i++ {
		maxNum = max(maxNum, nums[i])
	}
	for i := 1; i <= num; i++ {
		for j := 0; j < numsLen; j++ {
			for n := 0; n <= maxNum*num; n++ {
				if _, ok := result[n]; ok {
					if result[n][numsLen]+i <= num {
						tmpN := n + i*nums[j]
						result[tmpN] = make([]int, len(nums)+1)
						copy(result[tmpN], result[n])
						result[tmpN][j] += i
						result[tmpN][numsLen] += i
					}
				}
			}
		}
	}
	fmt.Println(result)
	fmt.Println(len(result) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
