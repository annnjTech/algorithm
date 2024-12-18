package model

type Node struct {
	Value int
	Next  *Node
}

type DoubleNode struct {
	Value int
	Last  *DoubleNode
	Next  *DoubleNode
}
