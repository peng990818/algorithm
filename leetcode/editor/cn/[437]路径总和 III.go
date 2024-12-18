//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。 
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
//输出：3
//解释：和等于 8 的路径有 3 条，如图所示。
// 
//
// 示例 2： 
//
// 
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：3
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [0,1000] 
// 
// -10⁹ <= Node.val <= 10⁹ 
// -1000 <= targetSum <= 1000 
// 
//
// Related Topics 树 深度优先搜索 二叉树 👍 1939 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func pathSum(root *TreeNode, targetSum int) int {
//     if root == nil {
//         return 0
//     }
//     var res int
//     process(&res, root, targetSum)
//     res += pathSum(root.Left, targetSum)
//     res += pathSum(root.Right, targetSum)
//     return res
// }
//
// func process(res *int, root *TreeNode, targetSum int) {
//     if root == nil {
//         return
//     }
//     if root.Val == targetSum {
//         *res++
//     }
//     process(res, root.Left, targetSum-root.Val)
//     process(res, root.Right, targetSum-root.Val)
// }

// func rootSum(root *TreeNode, targetSum int) (res int) {
// if root == nil {
// return
// }
// val := root.Val
// if val == targetSum {
// res++
// }
// res += rootSum(root.Left, targetSum-val)
// res += rootSum(root.Right, targetSum-val)
// return
// }
//
// func pathSum(root *TreeNode, targetSum int) int {
// if root == nil {
// return 0
// }
// res := rootSum(root, targetSum)
// res += pathSum(root.Left, targetSum)
// res += pathSum(root.Right, targetSum)
// return res
// }

func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    var res int
    process(&res, root, targetSum)
    res += pathSum(root.Left, targetSum)
    res += pathSum(root.Right, targetSum)
    return res
}

func process(res *int, root *TreeNode, target int) {
    if root == nil {
        return
    }
    if root.Val == target {
        *res++
    }
    process(res, root.Left, target-root.Val)
    process(res, root.Right, target-root.Val)
}
//leetcode submit region end(Prohibit modification and deletion)
