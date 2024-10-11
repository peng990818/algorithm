//给定一个二叉树，判断它是否是 平衡二叉树 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：true
// 
//
// 示例 2： 
// 
// 
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：true
// 
//
// 
//
// 提示： 
//
// 
// 树中的节点数在范围 [0, 5000] 内 
// -10⁴ <= Node.val <= 10⁴ 
// 
//
// Related Topics 树 深度优先搜索 二叉树 👍 1553 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, b := isBalanced2(root)
    return b
}

func isBalanced2(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    depthL, bL := isBalanced2(root.Left)
    depthR, bR := isBalanced2(root.Right)
    if !bL || !bR {
        return 0, false
    }
    d, max := 0, 0
    if depthL > depthR {
        d = depthL - depthR
        max = depthL
    } else {
        d = depthR - depthL
        max = depthR
    }
    if d <= 1 {
        return max+1, true
    }
    return max+1, false
}

//leetcode submit region end(Prohibit modification and deletion)
