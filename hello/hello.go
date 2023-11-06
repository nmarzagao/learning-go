package main
import (
	"fmt"
)

func main() {
	var name string = getName()
	printHello(name)

	countdown(10)

	list := &List{}
    list.insert(1)
    list.insert(2)
    list.insert(3)
    list.insert(4)

	tmp := list.head
    fmt.Println("Initial List: ")
    printList(tmp)

	tmp = list.head
    list.remove(2)
    fmt.Println("List after removing 2: ")
    printList(tmp)

	tmp = list.head
    list.remove(4)
    fmt.Println("List after removing 4: ")
    printList(tmp)
}