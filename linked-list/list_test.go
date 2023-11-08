package linked_list

import (
	"testing"
	"fmt"
)

func test_linked_list(t *testing.T) {
	list := &List{head: nil}

	list.insert(1)
	list.insert(2)
	list.insert(3) 
	
	if list.head.data != 1 {
		t.Errorf("Expected the first node to be '1', but got %v\n", list.head.data)
	} else if list.head.next.data != 2 {
		t.Errorf("Expected the second node to be '2', but got %v\n", list.head.next.data)
	} else if list.head.next.next.data != 3 {
		t.Errorf("Expected the third node to be '3', but got %v\n", list.head.next.next.data)
	} else {
		fmt.Println("Insert is working fine!")
	}

	list.remove(2)
	if list.head.next.data == 2 {
		t.Errorf("Expected th node to be removed, but got %v\n", list.head.next.data)
	} else {
		fmt.Println("Remove is working fine!")
	}
}