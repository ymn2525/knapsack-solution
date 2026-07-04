// 方針: 動的計画法で解く
// 参考: https://qiita.com/drken/items/a5e6fe22863b7992efdb
package main

import (
	"fmt"
	//"time"
)

type Item struct {
	size int
	value int
}

func main() {
	// ナップサックの容量
	W := 45

	// 品物の数
	var N int = 18

	// 品物の配列{容量, 値段}
	items := []Item{
		{W+1, -1}, // パディング
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
	// 動的計画法の実装
	
	// 解計算用テーブルの作成
	// Go言語の仕様により、初期化時の要素は全て0
	dp := make([][][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([][]int, W+1)
		for j := 0; j < W+1; j++ {
			dp[i][j] = make([]int, 1, N+1)
		}
	}

	for n := 1; n < N+1; n++ {
		for w := 1; w < W+1; w++ {
			if w >= items[n].size  {
				if dp[n-1][w][0] > dp[n-1][w-items[n].size][0] + items[n].value {
					dp[n][w] = append([]int{dp[n-1][w][0]}, dp[n-1][w][1:]...)
				}else{
					dp[n][w] = append([]int{dp[n-1][w-items[n].size][0] + items[n].value}, dp[n-1][w-items[n].size][1:]...)
					dp[n][w] = append(dp[n][w], n)
				}
			}else{
				dp[n][w] = append([]int{dp[n-1][w][0]}, dp[n-1][w][1:]...)
			}
		}
	}

	fmt.Printf("品物リスト : %v\n", dp[N][W][1:])
	fmt.Printf("最大価値 : %d\n", dp[N][W][0])
	fmt.Printf("容量 : %d\n", 
		func() int {
			var i int
			for i = 1; i < W+1; i++ {
				if dp[N][W][0]== dp[N][i][0] {
					return i
				}
			}
			return -1
		}(),
	)
}
