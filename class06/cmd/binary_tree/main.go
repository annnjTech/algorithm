package main

import (
	"algorithm/class06/binary_tree"
	"algorithm/class06/model"
	"fmt"
)

func main() {
	head := &model.TreeNode{
		Value: 1,
		Left:  nil,
		Right: nil,
	}
	head.Left = &model.TreeNode{
		Value: 2,
		Left:  nil,
		Right: nil,
	}
	head.Right = &model.TreeNode{
		Value: 3,
		Left:  nil,
		Right: nil,
	}
	head.Left.Left = &model.TreeNode{
		Value: 4,
		Left:  nil,
		Right: nil,
	}
	head.Left.Right = &model.TreeNode{
		Value: 5,
		Left:  nil,
		Right: nil,
	}
	head.Right.Left = &model.TreeNode{
		Value: 6,
		Left:  nil,
		Right: nil,
	}
	head.Right.Right = &model.TreeNode{
		Value: 7,
		Left:  nil,
		Right: nil,
	}

	fmt.Println("==========")
	binary_tree.Pre(head)
	fmt.Println("==========")
	binary_tree.In(head)
	fmt.Println("==========")
	binary_tree.Post(head)
	fmt.Println("==========")
}

/*
	      ①
        /    \
	   ②     ③
	  / \    / \
     ④  ⑤  ⑥  ⑦
1,2,4,4,4,2,5,5,5,2,1,3,6,6,6,3,7,7,7,3,1
先序遍历就是数字第一次出现的顺序，即1,2,4,5,3,6,7
中序遍历就是数字第二次出现的顺序，即4,2,5,1,6,3,7
后序遍历就是数字第三次出现的顺序，即4,5,2,6,7,3,1
*/
