package symmetric_tree

import "algorithm/class06/model"

// IsSymmetric 判断是否是对称二叉树（镜面树）
func IsSymmetric(root *model.TreeNode) bool {
	return IsMirror(root, root)
}

func IsMirror(h1, h2 *model.TreeNode) bool {
	if (h1 == nil) != (h2 == nil) {
		return false
	}
	if h1 == nil && h2 == nil {
		return true
	}
	return h1.Value == h2.Value && IsMirror(h1.Left, h2.Right) && IsMirror(h1.Right, h2.Left)
}
