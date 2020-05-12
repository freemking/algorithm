package main

import (
	"fmt"
)

/*
第一步要明确两点，「状态」和「选择」
第二步要明确 dp 数组的定义
第三步，根据「选择」，思考状态转移的逻辑
最后一步，把伪码翻译成代码，处理一些边界情况
*/

func main() {
	//最长增长子序列
	maxUpLength()
	//0-1背包问题
	zoreToOnePackage()
	//完全背包问题
	fullPackage()
	//子集背包问题
	fmt.Println(sonPackage())
	//正则表达式
	fmt.Println(isMatch("", ".*a*b*"))
}

func maxUpLength() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	var dp []int
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
}

func zoreToOnePackage() {
	n := 3
	w := 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}

	dp := make([][]int, n+1)

	dp[0] = append(dp[0], 0, 0, 0, 0, 0)
	for i := 1; i <= n; i++ {
		dp[i] = append(dp[i], 0, 0, 0, 0, 0)
		for j := 1; j <= w; j++ {
			if j > wt[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	fmt.Println(dp[n][w])
}

func fullPackage() {
	coins := []int{1, 2, 5}
	amount := 5
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	fmt.Println(dp[amount])
}

func sonPackage() bool {
	nums := []int{7, 8, 9, 10}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[sum]
}

func isMatch(text, pattern string) bool {
	if text == pattern {
		return true
	}
	if pattern == "" {
		return false
	}

	matchStatus := text != "" && (pattern[0] == '.' || text[0] == pattern[0])

	if len(pattern) > 1 && pattern[1] == '*' {
		return (matchStatus && isMatch(text[1:], pattern)) || isMatch(text, pattern[2:])
	}

	return matchStatus && isMatch(text[1:], pattern[1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
