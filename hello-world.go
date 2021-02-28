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
	if head == nil {
		return nil, nil
	}
	return head.next, head
}

var currentTicket int = 1

func getTicket() int {
	currentTicket = currentTicket + 1
	return currentTicket - 1
}

func printList(list *node) {
	if list == nil {
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

func retrieve(ticket int, alleyA *node, alleyB *node) (*node, *node) {
	var car *node = alleyA
	if alleyA == nil {
		return nil, nil
	}
	// search while the current car isn't the one we're looking for
	// and the car isn't the last in the alley.
	// for each car we're not interested in, insert it into allyB.
	for alleyA.data != ticket && alleyA.next != nil {
		alleyA, car = remove(alleyA)
		alleyB = insert(alleyB, car.data)
		fmt.Println("A")
		printList(alleyA)
		fmt.Println("B")
		printList(alleyB)
	}
	// if we've found the correct car, return it
	if alleyA.data == ticket {
		fmt.Println("Found ", ticket)
		fmt.Println("A")
		printList(alleyA)
		fmt.Println("B")
		printList(alleyB)
		alleyA, _ = remove(alleyA) // this effectively "drops" the item from the linked list
		fmt.Println("B")
		printList(alleyB)
		// move contents of alleyB back into alleyA
		for alleyB != nil {
			var putBack *node
			alleyB, putBack = remove(alleyB)
			alleyA = insert(alleyA, putBack.data)
		}
		printList(alleyA)
		return alleyA, car
	} else {
		// this is when the car wasn't found; move alleyB back into alleyA.
		// nil signifies the car wasn't found
		for alleyB != nil {
			alleyB, car = remove(alleyB)
			alleyA = insert(alleyA, car.data)
		}
		return alleyA, nil
	}
}

func main() {
	var alleyA *node
	var alleyB *node

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
			var ticket = getTicket()
			fmt.Println("Ticket no.=", ticket)
			alleyA = insert(alleyA, ticket)
		case 3:
			fmt.Println("case 3")
			fmt.Println("Which ticket number?")
			var ticket int
			fmt.Scanf("%d", &ticket)
			var node *node
			alleyA, node = retrieve(ticket, alleyA, alleyB)
			if node == nil {
				fmt.Println("ticket not found", ticket)
			} else {
				fmt.Println("removed node with ticket: ", node.data)
			}
		case 5: // this is a remove call
			var node *node
			alleyA, node = remove(alleyA)
			if node == nil {
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
