//给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。 
//
// 
//
// 示例 1： 
//
// 
// 输入：root = [1,null,2,3] 
// 
//
// 输出：[3,2,1] 
//
// 解释： 
//
// 
//
// 示例 2： 
//
// 
// 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9] 
// 
//
// 输出：[4,6,7,5,2,9,8,3,1] 
//
// 解释： 
//
// 
//
// 示例 3： 
//
// 
// 输入：root = [] 
// 
//
// 输出：[] 
//
// 示例 4： 
//
// 
// 输入：root = [1] 
// 
//
// 输出：[1] 
//
// 
//
// 提示： 
//
// 
// 树中节点的数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？ 
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1198 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func postorderTraversal(root *TreeNode) (res []int) {
//     // 递归实现
//     if root == nil {
//         return
//     }
//     res = append(res, postorderTraversal(root.Left)...)
//     res = append(res, postorderTraversal(root.Right)...)
//     res = append(res, root.Val)
//     return
//     // todo 循环实现
//     // todo morisi遍历实现
// }

func postorderTraversal(root *TreeNode) (res []int) {
if root == nil {
return nil
}
stack := []*TreeNode{}
stack = append(stack, root)
for len(stack) > 0 {
tmp := stack[len(stack)-1]
stack = stack[:len(stack)-1]
res = append(res, tmp.Val)
if tmp.Left != nil {
stack = append(stack, tmp.Left)
}
if tmp.Right != nil {
stack = append(stack, tmp.Right)
}
}
for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
res[i], res[j] = res[j], res[i]
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
