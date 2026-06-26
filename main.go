package main

import (
	"fmt"
	"math"
)

type Item struct {
	size int
	value int
}

func index_calc(num int) []int {
	bin := make([]int, 18)
	for i := 0; num > 0; i++ {
		bin[i] = num % 2
		num = num / 2
	}
	return bin
}


func main() {

	items := [18]Item{
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
	
	// 最終回答用変数の定義
	max_value := 0
	max_size := 0
	max_items := make([]int, 0, 18)

	// ループの上限回数
	max_roop := int(math.Pow(2, 18))

	for i := 0; i < max_roop; i++ {
		value := 0
		size := 0
		ans_items := make([]int, 0, 18)
		bin := index_calc(i)
		
		// iの値を2進数で扱えるようにし、1に該当する品物の価値の総和を求めるのをiを0~2^{18}まで繰り返す
		for j, b := range bin {
			if b == 1 {
				value += items[j].value
				size += items[j].size
				ans_items = append(ans_items, j)
			}
		}

		// 合計sizeが45以下　かつ　価値が現時点の最大値以上　ならば　最大値を上書きしてその品物の配列を記憶する
		if size <= 45 && value >= max_value {
			max_value = value
			max_size = size
			max_items = make([]int, len(ans_items))
			copy(max_items, ans_items)
		}
	}
	
	fmt.Printf("品物: [ %v", max_items[0]+1)
	for _, index := range max_items[1:] {
		fmt.Printf(", %v", index+1)
	}
	fmt.Println(" ]")
	fmt.Printf("最大価値: %v\n", max_value)
	fmt.Printf("容量: %v\n", max_size)

}
