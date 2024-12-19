package main

import (
	"algorithm/class06/binary_tree"
	"algorithm/class06/construct_binary_tree"
)

func main() {
	head := construct_binary_tree.BuildTreeFromPreorderInorder(
		[]int{1, 2, 4, 5, 3, 6, 7},
		[]int{4, 2, 5, 1, 6, 3, 7},
	)
	binary_tree.Post(head)
}
