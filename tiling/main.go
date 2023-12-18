package main

import "fmt"

func countTileWorkMethods(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	var n int

	fmt.Scanln(&n)

	fmt.Println(countTileWorkMethods(n))
}
