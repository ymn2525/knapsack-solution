// 方針: 動的計画法で解く
package main


func DynamicProgramming() (int, int, []int){
	// ナップサックの容量W, 品物の数N
	W, N := 45, 18

	// 品物の配列{容量, 値段}
	items := []Item{
		{0, 0}, // パディング
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
	dp_table := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp_table[i] = make([]int, W+1)
	}

	// dp表において、重さに対応した表。次のdp表を作る過程で変化していく。
	// wが1~Wの場合のitemリスト。1つのwに対してn個配列を用意しないでいい理由は、次のdp表作成の過程ではn=3のとき、n-2=1列目の情報はいらないため。
	selected := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		selected[i] = make([][]int, W+1)
		for j := 0; j < W+1; j++ {
			selected[i][j] = make([]int, N+1)
		}
	}
	last_selected := false

	for n := 1; n < N+1; n++ {
		for w := 1; w < W+1; w++ {
			if w >= items[n].size && dp_table[n-1][w] <= dp_table[n-1][w-items[n].size] + items[n].value  {
				dp_table[n][w] = dp_table[n-1][w-items[n].size] + items[n].value
				copy(selected[0][w], selected[1][w-items[n].size])
				selected[0][w][n] = n
			}else{
				dp_table[n][w] = dp_table[n-1][w]
				copy(selected[0][w], selected[1][w])
			}
			
		}
		last_selected = !last_selected
		selected[0], selected[1] = selected[1], selected[0]
	}
  
	var last int
	if last_selected {
		last = 1
	}else{
		last = 0
	}

	size := func() int {
			var i int
			for i = 1; i < W+1; i++ {
				if dp_table[N][W] == dp_table[N][i] {
					return i
				}
			}
			return -1
		}()
	return dp_table[N][W], size, selected[last][W] 
}
