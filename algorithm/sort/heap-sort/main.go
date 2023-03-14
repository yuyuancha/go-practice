package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{4, 5, 1, 6, 7, 9, 2, 8, 3}
	printTree(arr)
}

func printTree(arr []int) {
	if len(arr) == 0 {
		return
	}

	var levelLast = 0
	var currentLevel = 0

	for index := 0; index < len(arr)-1; index++ {
		fmt.Printf("%d ", arr[index])

		if levelLast == index {
			fmt.Printf("\n")
			currentLevel++
			levelLast += int(math.Pow(2, float64(currentLevel)))
		}
	}

	fmt.Printf("\n")
}
