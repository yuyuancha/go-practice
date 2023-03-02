package main

import "fmt"

func main() {
	// 建立一個大小為 4 的陣列
	var arr [4]int
	arr[0] = 10
	arr[1] = 4
	arr[2] = 8
	arr[3] = 9

	// 走訪此陣列
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Printf("\n")
}
