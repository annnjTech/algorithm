package max_depth

import (
	"algorithm/class06/model"
	"math"
)

// MaxDepth 计算二叉树的最大深度
/*
        ①
       /  \
      ②   ③
            \
             ④
思路：
1) root为①节点，不为空
2) 递归①的左子树，root为②节点，不为空
3) 递归②的左子树，root为nil，返回0
4) 递归②的右子树，root为nil，返回0
5) max(0, 0)+1=1，返回1，回到步骤2)处
6) 递归①的右子树，root为③节点，不为空
7) 递归③的左子树，root为nil，返回0
8) 递归③的右子树，root为④节点，不为空
9) 递归④的左子树，root为nil，返回0
10) 递归④的右子树，root为nil，返回0
11) max(0,0)+1=1，返回1，回到步骤8)处
12) max(步骤7的0, 步骤11的1) + 1 = 2，回到步骤6)处
13) max(步骤5的1，步骤12的2) + 1 = 3，回到步骤2)处，即回到步骤1)的语句return int(math.Max(1.0, 2.0)) + 1 = 3，所以结果为3层
*/
func MaxDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max((float64)(MaxDepth(root.Left)), (float64)(MaxDepth(root.Right)))) + 1
}
