package main

import (
	"fmt"
	"strconv"
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
	//编辑距离
	fmt.Println(editDistance("rad", "apple"))
	//高楼扔鸡蛋
	memo := map[string]int{}
	fmt.Println(eggTallBuildingLinear(100, 2, memo))    //线性搜索
	fmt.Println(eggTallBuildingDichotomy(100, 2, memo)) //二分搜索
	//戳气球
	fmt.Println(pokeBalloon())
	//最长公共子序列
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	//最长回文子串
	fmt.Println(longestPalindromeSubseq("bbaa"))
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

func editDistance(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if s1[0] == s2[0] {
		return editDistance(s1[1:], s2[1:])
	} else {
		insert := editDistance(s1[0:], s2[1:]) + 1
		delete := editDistance(s1[1:], s2[0:]) + 1
		replace := editDistance(s1[1:], s2[1:]) + 1
		return min(min(insert, delete), replace)
	}
}

func eggTallBuildingLinear(height, num int, memo map[string]int) int {
	if height == 0 {
		return 0
	}
	if num == 1 {
		return height
	}
	key := strconv.Itoa(height) + "-" + strconv.Itoa(num)
	if val, ok := memo[key]; ok {
		return val
	}
	res := int(^uint(0) >> 1)
	for i := 1; i <= height; i++ {
		res = min(res, max(eggTallBuildingLinear(i-1, num-1, memo), eggTallBuildingLinear(height-i, num, memo))+1)
	}
	memo[key] = res
	return res
}

func eggTallBuildingDichotomy(height, num int, memo map[string]int) int {
	if height == 0 {
		return 0
	}
	if num == 1 {
		return height
	}
	key := strconv.Itoa(height) + "-" + strconv.Itoa(num)
	if val, ok := memo[key]; ok {
		return val
	}
	res := int(^uint(0) >> 1)
	low, high := 1, height
	for low <= high {
		mid := (low + high) / 2
		broken := eggTallBuildingDichotomy(mid-1, num-1, memo)
		no_broken := eggTallBuildingDichotomy(height-mid, num, memo)
		if broken > no_broken {
			high = mid - 1
			res = min(res, broken+1)
		} else {
			low = mid
			res = min(res, no_broken+1)
		}
	}
	memo[key] = res
	return res
}

func pokeBalloon() int {
	nums := []int{3, 1, 5, 8}
	n := len(nums)
	// 添加两侧的虚拟气球
	points := make([]int, n+2)
	points[0] = 1
	points[n+1] = 1
	copy(points[1:n+1], nums[0:])
	// base case 已经都被初始化为 0
	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
	}
	// 开始状态转移
	// i 应该从下往上
	for i := n; i >= 0; i-- {
		// j 应该从左往右
		for j := i + 1; j < n+2; j++ {
			// 最后戳破的气球是哪个？
			for k := i + 1; k < j; k++ {
				// 择优做选择
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}
	return dp[0][n+1]
}

func longestCommonSubsequence(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	var s string
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				s += string(s1[i-1])
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Println(s)
	return dp[m][n]
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[n-1][n-1] = 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[0][n-1]
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
