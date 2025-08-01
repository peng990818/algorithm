//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。 
//
// 有效 二叉搜索树定义如下： 
//
// 
// 节点的左子树只包含 小于 当前节点的数。 
// 节点的右子树只包含 大于 当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
// 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [2,1,3]
//输出：true
// 
//
// 示例 2： 
// 
// 
//输入：root = [5,1,4,null,null,3,6]
//输出：false
//解释：根节点的值是 5 ，但是右子节点的值是 4 。
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目范围在[1, 10⁴] 内 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 2439 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isValidBST(root *TreeNode) bool {
//     return helper(root, math.MinInt64, math.MaxInt64)
// }
//
// func helper(root *TreeNode, lower, upper int) bool {
//     if root == nil {
//         return true
//     }
//     // 所有的值都应该在一个范围
//     if root.Val <= lower || root.Val >= upper {
//         return false
//     }
//     return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
// }


// 左树值的范围均小于root，右树值的范围均大于root
// func helper(root *TreeNode, lower, upper int) bool {
//     if root == nil {
//         return true
//     }
//     if root.Val <= lower || root.Val >= upper {
//         return false
//     }
//     return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
// }
//
// func isValidBST(root *TreeNode) bool {
//     return helper(root, math.MinInt64, math.MaxInt64)
// }

// 中序遍历方式的递归
// var pre *TreeNode
// func isValidBST(root *TreeNode) bool {
// if root == nil {
// return true
// }
// bl := isValidBST(root.Left)
// if pre != nil && pre.Val >= root.Val {
// return false
// }
// pre = root
// br := isValidBST(root.Right)
// return bl && br
// }
//leetcode submit region end(Prohibit modification and deletion)
