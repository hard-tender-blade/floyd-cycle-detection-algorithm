package main

func main() {
	ShowCase20()
}

func ShowCase20() {
	ll := LinkedList{}
	ll.AddNode(0)
	ll.AddNode(1)
	ll.AddNode(2)
	ll.AddNode(3)
	ll.AddNode(4)
	ll.AddNode(5)
	ll.AddNode(6)
	ll.AddNode(7)
	ll.AddNode(8)
	ll.AddNode(9)
	node10 := ll.AddNode(10)
	ll.AddNode(11)
	ll.AddNode(12)
	ll.AddNode(13)
	ll.AddNode(14)
	ll.AddNode(15)
	ll.AddNode(16)
	ll.AddNode(17)
	ll.AddNode(18)
	ll.AddNode(19)
	node20 := ll.AddNode(20)

	ll.SetNodeAfter(node10, node20)

	loopStartNode, found := ll.CycleDetection()
	if found {
		println("Loop detected, start node is: ", loopStartNode.Data)
	} else {
		println("Loop not detected")
		ll.PrintNodes()
	}
}
