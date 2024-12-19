package same_tree

import "algorithm/class06/model"

// IsSameTree 判断两棵二叉树是否相同
func IsSameTree(p, q *model.TreeNode) bool {
	if (p == nil) != (q == nil) { // 异或运算符，当两者不同时返回true
		return false
	}
	if p == nil && q == nil {
		return true
	}
	return p.Value == q.Value && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
