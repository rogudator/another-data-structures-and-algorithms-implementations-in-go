package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint)
		toPrint = toPrint.Next
		l.length--
	}
	fmt.Println()
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.Data == value {
		l.head = l.head.Next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.Next.Data != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
}