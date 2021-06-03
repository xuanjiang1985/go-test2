package main

import "fmt"

type listNode struct {
	val  int
	next *listNode
}

func main() {
	var head = &listNode{
		val: 1,
		next: &listNode{
			val: 2,
			next: &listNode{
				val: 3,
				next: &listNode{
					val: 4,
					next: &listNode{
						val:  5,
						next: nil,
					},
				},
			},
		},
	}
	newList := reverselist(head)
	printList(newList)
}

func reverselist(head *listNode) *listNode {
	var pre *listNode
	cur := head
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func printList(head *listNode) {

	for head != nil {
		fmt.Printf("%#v\n", head)
		head = head.next
	}
}
