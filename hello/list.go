package main
import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

//func (l *List) isEmpty() bool {
//	return (l.head == nil)
//}

func (l *List) insert(value int) {
	newNode := &Node{data: value}

    if l.head == nil {
        l.head = newNode
        return
    }

    curr := l.head
    for curr.next != nil {
        curr = curr.next
    }

    curr.next = newNode
}


func (l *List) remove(value int) {
	if l.head == nil {
        return
    }

    if l.head.data == value {
        l.head = l.head.next
        return
    }

    curr := l.head
    for curr.next != nil && curr.next.data != value {
        curr = curr.next
    }

    if curr.next != nil {
        curr.next = curr.next.next
    }
}

func printList(l *Node) {
	if l == nil {
		fmt.Println("nil")
		fmt.Println("");
		return
	} else {
		fmt.Printf("%d -> ", l.data)
		printList(l.next)
	}
}