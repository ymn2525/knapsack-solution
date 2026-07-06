package main

import "fmt"

func itemPrint(itemList []int) {
	fmt.Println("品物の組合せ : ")
	for _, item := range itemList {
		if item != 0 {
			fmt.Printf("%3d, ", item)
		}
	}
	fmt.Println()
}



func main()  {
	fmt.Println("============総当り法による実行============")
	value, size, selected_item := BruteForce()
	itemPrint(selected_item)
	fmt.Printf("最大価値: %v\n", value)
	fmt.Printf("容量: %v\n", size)


	fmt.Println("============動的計画法による実行============")
	value, size, selected_item = DynamicProgramming()
	itemPrint(selected_item)
	fmt.Printf("最大価値: %v\n", value)
	fmt.Printf("容量: %v\n", size)
}
