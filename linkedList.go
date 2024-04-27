package main

import "fmt"

type Node struct {
	Data int
	Next *Node
	Perv *Node
}

type LinkedList struct {
	Head      *Node
	Tail      *Node
	NodeCount int
}

func (ll *LinkedList) AddNode(data int) *Node {
	node := Node{Data: data}
	if ll.Head == nil {
		ll.Head = &node
		ll.Tail = &node
	} else {
		ll.Tail.Next = &node
		node.Perv = ll.Tail
		ll.Tail = &node
	}
	ll.NodeCount++
	return &node
}

func (ll *LinkedList) SetNodeAfter(node *Node, afterNode *Node) {
	afterNode.Next = node
}

func (ll *LinkedList) PrintNodes() {
	node := ll.Head
	for node != nil {
		fmt.Printf(" %d ->", node.Data)
		node = node.Next
	}
}

func (ll *LinkedList) CycleDetection() (*Node, bool) {
	slow := ll.Head
	fast := ll.Head
	found := false

	for fast != nil && fast.Next != nil {
		if !found {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}

		if found && slow == fast {
			return slow, true
		}
		if slow == fast {
			found = true
			slow = ll.Head
		}
	}
	return nil, false
}
