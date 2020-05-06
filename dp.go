package main

import (
	"fmt"
)

```
第一步要明确两点，「状态」和「选择」
第二步要明确 dp 数组的定义
第三步，根据「选择」，思考状态转移的逻辑
最后一步，把伪码翻译成代码，处理一些边界情况
```

func main() {
	//最长增长子序列
	maxUpLength()
	//0-1背包问题
	zoreToOnePackage()
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
	N = 3, W = 4
	wt = [2, 1, 3]
	val = [4, 2, 3]
	
	dp[n][w] = max(dp[n-1][w],dp[n-1][w-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
