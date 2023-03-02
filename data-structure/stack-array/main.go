package main

import (
	"fmt"
)

var top = -1
var arr [4]*int

func main() {
	// 透過陣列(arr [4]int)來實作關於 stack(堆疊)的特性——FILO(先進後出)。

	// 先依序推入 6、3、7
	push(6)
	push(3)
	push(7)

	// 取出最後一個值
	pop()

	// 再推入 8
	push(8)

	// 再取出最後一個值四次
	pop()
	pop()
	pop()
	pop()
}

// printArray 打印陣列
func printArray() {
	fmt.Println("top:", top)
	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			fmt.Printf("%d ", *arr[i])
		} else {
			fmt.Printf("nil ")
		}
	}
	fmt.Printf("\n")
}

// push 將值推入陣列中
func push(value int) {
	fmt.Println("將", value, "push 到堆疊中")

	if top == len(arr)-1 {
		return
	}

	top++
	arr[top] = &value

	printArray()
}

// pop 將最後放入得值取出
func pop() {
	fmt.Printf("將堆疊最後的值 pop 出來: ")

	if top == -1 {
		fmt.Printf("empty array\n")

		return
	}

	fmt.Printf("%d\n", *arr[top])
	arr[top] = nil
	top--

	printArray()
}
