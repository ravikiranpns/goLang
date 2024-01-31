package main

import "fmt"

func longestCommonSubsequence(t1 string, t2 string) int {
	dp := make([][]int, len(t1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t2)+1)
	}

	for i := len(t1) - 1; i >= 0; i-- {
		for j := len(t2) - 1; j >= 0; j-- {
			if t1[i] == t2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	t1 := "ab"
	t2 := "abc"

	fmt.Println(longestCommonSubsequence(t1, t2))
}
