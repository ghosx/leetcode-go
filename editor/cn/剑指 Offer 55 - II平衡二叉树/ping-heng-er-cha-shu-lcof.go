/**
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



 示例 1:

 给定二叉树 [3,9,20,null,null,15,7]


    3
   / \
  9  20
    /  \
   15   7

 返回 true 。

示例 2:

 给定二叉树 [1,2,2,3,3,null,null,4,4]


       1
      / \
     2   2
    / \
   3   3
  / \
 4   4


 返回 false 。



 限制：


 0 <= 树的结点个数 <= 10000


 注意：本题与主站 110 题相同：https://leetcode-cn.com/problems/balanced-binary-tree/


 Related Topics 树 深度优先搜索 二叉树 👍 246 👎 0

*/

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
