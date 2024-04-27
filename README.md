# Implementation and info about Floyd cycle detection algorithm

## Usage

Cycle detection algorithm is used in two cases

- We need to find the loop in liked list
- We need to find a duplicate in array of numbers where each number is less then array length

## Concept

Time complexity O(n) \
Space complexity O(1)

Keys:

- Slow and fast pointers go through linked list
- Slow moves one time (+1), fast moves to times (+2)
- If they met here is a loop, slow pointer move to the start of list
- When they meet second time on same node meant it is cycle start

Math proof:

```
x - path from list start node to cycle start node
y - path from cycle start node to first pointers contact node
z - path from first contact node to loop start node
l = y + z, it is cycle length

s = x + y + c1 * l, it is fast pointer path, c1 is cycle count constant can be 0 or bigger
f = x + y + c2 * l, it is fast pointer path, c2 is slow pointer cycle count constant

Proof:
f = 2s

x + y + c1 * l = 2(x + y + c2 * l) = 2x + 2y + 2(c2 * l)
c1 * l - 2c2 * l == 2x + 2y - x -y
l * (c1 - 2c2) = x + y
x = l * (c1 - 2c2) + y

x is equal to sudden amount of loops minus y

x = l * (c1 - 2c2) + l - y
x = l * (c3) + l - y
x = l * (c3) + y + z - y
x = l * (c3) + z

In simplest case we dont have any cycles, so l * (c3) will equal 0

x = 0 + z
x = z

```

### Historical context

Floyd's cycle-finding algorithm emerged in the 1960s, a time of growing interest in efficient algorithms for data structures like linked lists. It offered a fast way to detect loops (cycles) within these structures, becoming a popular choice for its space and time efficiency. While there are newer approaches, Floyd's algorithm remains a fundamental technique.

### Class diagram

```
classDiagram
class Node {
    Data: int
    Next: Node*
    Perv: Node*
}

class LinkedList {
    Head: Node*
    Tail: Node*
    NodeCount: int
}

LinkedList --> Node : AddNode(data int)
LinkedList --> Node : SetNodeAfter(node *Node, afterNode *Node)
LinkedList --> void : PrintNodes()
LinkedList --> Node*, bool : CycleDetection()

```
