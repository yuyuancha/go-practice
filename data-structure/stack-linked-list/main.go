package main

import (
	"fmt"
)

// node 鏈結串列的節點
type node struct {
	value int   // 此節點之值
	next  *node // 紀錄下一個節點位置
}

var list *node = nil

func main() {
	// 透過鏈結串列來實作關於 stack(堆疊)的特性——FILO(先進後出)。

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

// printLinkedList 走訪鏈結串列
func printLinkedList(n node) {
	fmt.Printf("%d->", n.value)
	if n.next == nil {
		fmt.Print("nil\n")

		return
	}

	printLinkedList(*n.next)
}

// push 將值推入陣列中
func push(value int) {
	fmt.Println("將", value, "push 到堆疊中")

	list = &node{
		value: value,
		next:  list,
	}

	printLinkedList(*list)
}

// pop 將最後放入得值取出
func pop() {
	fmt.Printf("將堆疊最後的值 pop 出來: ")

	if list == nil {
		fmt.Printf("empty linked list\n")

		return
	}

	fmt.Printf("%d\n", list.value)
	list = list.next

	if list != nil {
		printLinkedList(*list)
	} else {
		fmt.Println("nil")
	}
}
