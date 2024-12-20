package path_sum

import "algorithm/class07/model"

var isSum bool

/*
【路径就是从头节点到叶子节点的路径，路径的和就是路径上所有节点值的和。】
问题：给定一个二叉树的根节点 root 和一个整数 sum ，判断是否存在一条从根节点到叶子节点的路径，使得路径上的节点值之和等于 sum 。
*/
func hasPathSum(root *model.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	isSum = false // 这里为啥要重新赋值为false？当多次调用hasPathSum时，之前的isSum值会保留下来，导致结果错误
	process(root, 0, sum)
	return isSum
}

func process(x *model.TreeNode, preSum, sum int) {
	if x.Left == nil && x.Right == nil {
		if x.Value+preSum == sum {
			isSum = true
		}
		return
	}
	if x.Left != nil {
		process(x.Left, preSum+x.Value, sum)
	}
	if x.Right != nil {
		process(x.Right, preSum+x.Value, sum)
	}
}
