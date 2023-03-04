package main

import "fmt"

func main() {
	var target = []int{3, 5, 4, 7, 2, 8, 6, 9, 1}

	fmt.Println("排序前的陣列:", target)

	var result = sortByBubble(target)

	fmt.Println("排序完的陣列:", result)
}

// sortByBubble 透過氣泡排序做小到大排序
func sortByBubble(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		var flag = false

		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				flag = true
			}
		}

		// 如果不需要排序了，則不用繼續執行迴圈
		if !flag {
			break
		}

		fmt.Printf("i=%d迴圈結果: %v\n", i, input)
	}

	return input
}
