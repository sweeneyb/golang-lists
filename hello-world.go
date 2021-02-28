package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func insert(head *node, data int) *node {
	n := &node{data: data}

	if head == nil {
		return n
	} else {
		n.next = head
		return n
	}
}

func remove(head *node) (*node, *node) {
  if(head == nil) {
    return nil, nil
  }
  return head.next, head
}

var currentTicket int = 1

func getTicket() int {
  currentTicket = currentTicket+1
  return currentTicket-1
}

func printList(list *node) {
  if(list == nil ) {
    fmt.Println("list is empty")
    return
  }
  // fmt.Println("%v,", list.data)
  var item = list
  fmt.Println(item.data)
  for item.next != nil {
    fmt.Println(item.next.data)
    item = item.next
  }

}

func main() {
  var alleyA *node

	fmt.Println("Select one of the following choices:")
	fmt.Println("1. d or D (To display the alley contents)")
	fmt.Println("2. p or P (To park a car)")
	fmt.Println("3. r or R (To retrieve a car)")
	fmt.Println("4. q or Q (To quit the program)")
	var choice int
	for choice != 4 {
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
      fmt.Println("Alley A:")
      printList(alleyA)
    case 2:
      var ticket = getTicket();
      fmt.Println("Ticket no.=", ticket)
      alleyA = insert(alleyA, ticket)
		case 3:
      fmt.Println("case 3")
    case 5:
      var node *node
      alleyA, node = remove(alleyA)
      if(node == nil) {
        fmt.Println("list was empty")
      } else {
        fmt.Println("removing: ", node.data)
      }
		case 4:
			fmt.Println("Thanks and goodbye!")
			// Exit successfully
			break
		}
	}
}
