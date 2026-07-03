// 方針: 動的計画法で解く
// 参考: https://qiita.com/drken/items/a5e6fe22863b7992efdb
package main

import (
	"fmt"
	//"time"
)

func main() {
	// ナップサックの容量
	W := 45

	// 品物の数
	var n int = 18

	// 品物の配列{容量, 値段}
	items := [][]int{
		{4, 6},
		{8, 12},
		{3, 4},
		{5, 3},
		{9, 7},
		{2, 1},
		{3, 3},
		{1, 2},
		{5, 7},
		{2, 3},
		{4, 4},
		{2, 2},
		{7, 10},
		{10, 13},
		{3, 5},
		{13, 16},
		{11, 14},
		{8, 9},
	}
	
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < n; i++ {
		for w := 0; w <= W; w++ {
			if w >= items[i][0] {
				dp[i+1][w] = max(dp[i][w-items[i][0]] + items[i][1], dp[i][w])
			}else{
				dp[i+1][w] = dp[i][w]
			}
		}
	}
	fmt.Println(dp[n][W])

}
