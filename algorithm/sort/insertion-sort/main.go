package main

import "fmt"

func main() {
	var target = []int{3, 5, 4, 7, 2, 8, 6, 9, 1}

	fmt.Println("排序前的陣列:", target)

	var result = sortByInsertion(target)

	fmt.Println("排序完的陣列:", result)
}

// sortByInsertion 透過插入排序做小到大排序
func sortByInsertion(input []int) []int {
	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}

		fmt.Printf("i=%d迴圈結果: %v\n", i, input)
	}

	return input
}
