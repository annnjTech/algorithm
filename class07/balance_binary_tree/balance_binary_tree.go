package balance_binary_tree

import (
	"algorithm/class07/model"
	"math"
)

type Info struct {
	IsBalanced bool
	Height     int
}

// IsBalanced 验证是否为平衡二叉树
// 平衡二叉树：左右子树的高度差的绝对值不超过1
func IsBalanced(root *model.TreeNode) bool {
	return process(root).IsBalanced
}

func process(x *model.TreeNode) *Info {
	if x == nil {
		return &Info{true, 0}
	}
	info := new(Info)
	leftInfo := process(x.Left)
	rightInfo := process(x.Right)
	info.IsBalanced = leftInfo.IsBalanced && rightInfo.IsBalanced && math.Abs((float64)(leftInfo.Height-rightInfo.Height)) < 2
	info.Height = 1 + int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height)))
	return info
}
