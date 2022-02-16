package main

import "fmt"

func main() {
	for n := 0; n < 10; n++ {
		list := getList(n)
		fmt.Println(swapPairs(list))
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) String() string {
	vals := make([]int, 0)
	for list != nil {
		vals = append(vals, list.Val)
		list = list.Next
	}
	return fmt.Sprint(vals)
}

func getList(n int) (head *ListNode) {
	var curr *ListNode

	for i := 0; i < n; i++ {
		node := ListNode{i, nil}
		if i == 0 {
			head = &node
			curr = &node
		} else {
			curr.Next = &node
			curr = &node
		}
	}

	return
}

func swapPairs(head *ListNode) *ListNode {
	// if len(list) <= 1
	if head == nil || head.Next == nil {
		return head
	}

	// len(list) >= 2
	newHead := head.Next
	i, j := head, head.Next
	for i != nil && j != nil {
		k := j.Next
		j.Next = i
		if k != nil && k.Next != nil {
			i.Next = k.Next

			i = k
			j = i.Next
		} else {
			i.Next = k
			break
		}
	}

	return newHead
}
