package binary_search_tree

import (
	"algorithm/class07/model"
	"math"
)

type Info struct {
	IsBST bool
	Max   int
	Min   int
}

// ValidBSTProcess
/*
 * 验证是否为二叉搜索树
 */
func ValidBSTProcess(root *model.TreeNode) *Info {
	if root == nil {
		return nil
	}
	info := new(Info)
	leftInfo := ValidBSTProcess(root.Left)
	rightInfo := ValidBSTProcess(root.Right)
	// 整棵树的最大值和最小值
	var tmpMax, tmpMin int
	tmpMax = root.Value
	tmpMin = root.Value
	if leftInfo != nil {
		tmpMax = int(math.Max((float64)(tmpMax), (float64)(leftInfo.Max)))
		tmpMin = int(math.Min((float64)(tmpMin), (float64)(leftInfo.Min)))
	}
	if rightInfo != nil {
		tmpMax = int(math.Max((float64)(tmpMax), (float64)(rightInfo.Max)))
		tmpMin = int(math.Min((float64)(tmpMin), (float64)(rightInfo.Min)))
	}

	var isBST bool = true
	if leftInfo != nil && !leftInfo.IsBST {
		isBST = false
	}
	if rightInfo != nil && !rightInfo.IsBST {
		isBST = false
	}
	if (leftInfo != nil && leftInfo.Max >= root.Value) || (rightInfo != nil && rightInfo.Min <= root.Value) {
		isBST = false
	}

	info.IsBST = isBST
	info.Max = tmpMax
	info.Min = tmpMin
	return info
}
