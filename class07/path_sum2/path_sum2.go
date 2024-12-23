package path_sum2

import "algorithm/class07/model"

// CollectPathSum
/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
*/
func CollectPathSum(root *model.TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var path []int
	process(root, &path, 0, sum, &res)
	return res
}

func process(node *model.TreeNode, path *[]int, preSum, sum int, res *[][]int) {
	// node为叶子节点
	if node.Left == nil && node.Right == nil {
		if preSum+node.Value == sum {
			*path = append(*path, node.Value)
			length := len(*path)
			copyPath := make([]int, length)
			copy(copyPath, *path)
			*res = append(*res, copyPath)
			*path = (*path)[:length-1]
		}
		return
	}
	// node为非叶子节点
	*path = append(*path, node.Value)
	preSum += node.Value
	if node.Left != nil {
		process(node.Left, path, preSum, sum, res)
	}
	if node.Right != nil {
		process(node.Right, path, preSum, sum, res)
	}
	length := len(*path)
	*path = (*path)[:length-1]
	return
}
