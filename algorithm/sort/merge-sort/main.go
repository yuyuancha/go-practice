package main

import "fmt"

func main() {
	var target = []int{3, 5, 4, 7, 2, 8, 6, 9, 1}

	fmt.Println("排序前的陣列:", target)

	var result = sortByMerge(target)

	fmt.Println("排序完的陣列:", result)
}

// sortByMerge 透過合併排序做小到大排序
func sortByMerge(input []int) []int {
	if len(input) <= 1 {
		// 剩下最後一個元素時，就開使返回
		return input
	}

	var middleIndex = len(input) / 2

	// 分別對左右兩側做遞迴合併排序
	var left, right = sortByMerge(input[:middleIndex]), sortByMerge(input[middleIndex:])
	var result []int
	var rightIndex = 0

	for i := 0; i < len(left); i++ {
		for j := rightIndex; j < len(right); j++ {
			if left[i] > right[j] {
				// 右邊的數字較小，就合併進去
				rightIndex = j + 1
				result = append(result, right[j])
			} else {
				// 左邊的數字較小，就合併進去並開始比較左邊下一個數
				result = append(result, left[i])
				break
			}
		}

		// 當右邊的數都放完，則將左邊剩餘的數合併進去
		if len(right) == rightIndex {
			result = append(result, left[i:]...)
			break
		}
	}

	// 檢查右邊的數有沒有放完，沒有就合併進去
	if len(right) > rightIndex {
		result = append(result, right[rightIndex:]...)
	}

	fmt.Println(result)

	return result
}
