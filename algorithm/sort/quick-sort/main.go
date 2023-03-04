package main

import "fmt"

func main() {
	var target = []int{3, 5, 4, 7, 2, 8, 6, 9, 1}

	fmt.Println("排序前的陣列:", target)

	var result = sortByQuick(target)

	fmt.Println("排序完的陣列:", result)
}

// sortByQuick 透過快速排序做小到大排序
func sortByQuick(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	var less, greater, pivotSlice []int
	var pivotIndex = len(input) / 2
	var pivot = input[pivotIndex]

	for i := 0; i < len(input); i++ {
		if i == pivotIndex {
			continue
		}

		if input[i] >= pivot {
			greater = append(greater, input[i])
		} else {
			less = append(less, input[i])
		}
	}

	pivotSlice = append(pivotSlice, pivot)

	var result = append(sortByQuick(less), pivotSlice...)
	result = append(result, sortByQuick(greater)...)

	fmt.Println("此次排序結果:", result)

	return result
}
