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
		{W+1, -1},
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
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for n := 1; n < N+1; n++ {
		for w := 1; w < W+1; w++ {
			if w >= items[n].size  {
				fmt.Println(dp[n-1][1:])
				fmt.Println(dp[n][1:])
				fmt.Printf("品物%d%v, 容量上限%dのとき、w >= items[n].size = %tです\n", n, items[n], w, w >= items[n].size)
				fmt.Printf("1マス上の価値は%d, 品物%dを加えた価値はdp[%d][%d] + %d = %dです\n", dp[n-1][w], n, n-1, w-items[n].size,items[n].value, dp[n-1][w-items[n].size]+items[n].value)
				fmt.Println()
				dp[n][w] = max(dp[n-1][w], dp[n-1][w-items[n].size]+items[n].value)
			}else{
				dp[n][w] = dp[n-1][w]
			}
		}
	}


	for i, list := range dp[1:] {
		fmt.Printf("%2d : ", i+1)
		for _, l := range list {
			fmt.Printf("%3d", l)
		}
		fmt.Println()
	}
}
