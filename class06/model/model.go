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

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}
