package binary_tree_level_order_traversal

import "algorithm/class07/model"

// LevelOrderBottom
/*
二叉树的层序遍历，从底层开始遍历
给你二叉树的根节点root，返回其节点值自底向上的层序遍历（即按从叶子节点所在层到根节点所在层的顺序，逐层从左向右遍历）
*/
func LevelOrderBottom(root *model.TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var queue []*model.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		var curAns []int
		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]
			curAns = append(curAns, curNode.Value)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		ans = append(ans, curAns)
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}
