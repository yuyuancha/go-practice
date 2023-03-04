package main

import "fmt"

func main() {
	var target = []int{3, 5, 4, 7, 2, 8, 6, 9, 1}

	fmt.Println("排序前的陣列:", target)

	var result = sortBySelection(target)

	fmt.Println("排序完的陣列:", result)
}

// sortBySelection 透過選擇排序做小到大排序
func sortBySelection(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		var minIndex = i

		for j := i + 1; j < len(input); j++ {
			if input[minIndex] > input[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			input[i], input[minIndex] = input[minIndex], input[i]
		}
	}

	return input
}
