package main

import "fmt"

// node 鏈結串列的節點
type node struct {
	value int   // 此節點之值
	next  *node // 紀錄下一個節點位置
}

func main() {
	// 建立一個 10 -> 4 -> 8 ->nil 的鏈結串列
	var list = node{
		value: 10,
		next: &node{
			value: 4,
			next: &node{
				value: 8,
				next:  nil,
			},
		},
	}

	printLinkedList(list)
}

// printLinkedList 走訪鏈結串列
func printLinkedList(list node) {
	fmt.Printf("%d->", list.value)
	if list.next == nil {
		fmt.Print("nil\n")

		return
	}

	printLinkedList(*list.next)
}
