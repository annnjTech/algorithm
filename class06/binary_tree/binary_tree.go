package binary_tree

import (
	"algorithm/class06/model"
	"fmt"
)

func find(head *model.TreeNode) {
	if head == nil {
		return
	}
	// 1 在此打印时，属于先序遍历：头左右
	find(head.Left)
	// 2 在此打印时，属于中序遍历：左头右
	find(head.Right)
	// 3 在此打印时，属于后序遍历：左右头
}

// Pre 先序遍历
func Pre(head *model.TreeNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Value)
	Pre(head.Left)
	Pre(head.Right)
}

// In 中序遍历
func In(head *model.TreeNode) {
	if head == nil {
		return
	}
	In(head.Left)
	fmt.Println(head.Value)
	In(head.Right)
}

// Post 后序遍历
func Post(head *model.TreeNode) {
	if head == nil {
		return
	}
	Post(head.Left)
	Post(head.Right)
	fmt.Println(head.Value)
}
