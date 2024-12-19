package construct_binary_tree

import "algorithm/class06/model"

func BuildTreeFromPreorderInorder(pre, in []int) *model.TreeNode {
	if pre == nil || in == nil || len(pre) != len(in) {
		return nil
	}
	return f(pre, 0, len(pre)-1, in, 0, len(in)-1)
}

func f(pre []int, l1, r1 int, in []int, l2, r2 int) *model.TreeNode {
	if l1 > r1 {
		return nil
	}

	head := &model.TreeNode{Value: pre[l1]}
	if l1 == r1 {
		return head
	}
	var findIndex = l2
	for in[findIndex] != pre[l1] {
		findIndex++
	}
	head.Left = f(pre, l1+1, l1+findIndex-l2, in, l2, findIndex-1)
	head.Right = f(pre, l1+findIndex-l2+1, r1, in, findIndex+1, r2)
	return head
}