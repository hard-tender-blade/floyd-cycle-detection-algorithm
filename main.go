package main

func main() {
	ll := LinkedList{}
	ll.AddNode(0)
	node1 := ll.AddNode(1)
	ll.AddNode(2)
	node3 := ll.AddNode(3)

	println(node1, node3)
	ll.SetNodeAfter(node1, node3)

	loopStartNode, found := ll.CycleDetection()
	if found {
		println("Loop detected, start node is: ", loopStartNode.Data)
	} else {
		println("Loop not detected")
		ll.PrintNodes()
	}

}
