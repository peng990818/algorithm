//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。 
//
// 假设二叉树中至少有一个节点。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: root = [2,1,3]
//输出: 1
// 
//
// 示例 2: 
//
// 
//
// 
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [1,10⁴] 
// 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 631 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func findBottomLeftValue(root *TreeNode) int {
// if root == nil {
// return 0
// }
// queue := []*TreeNode{root}
// var res int
// for len(queue) > 0 {
// size := len(queue)
// res = queue[0].Val
// for i := 0; i < size; i++ {
// if queue[i].Left != nil {
// queue = append(queue, queue[i].Left)
// }
// if queue[i].Right != nil {
// queue = append(queue, queue[i].Right)
// }
// }
// queue = queue[size:]
// }
// return res
// }

// 回溯
func findBottomLeftValue(root *TreeNode) int {
if root == nil {
return 0
}
maxDepth := math.MinInt32
result := 0
process(&result, &maxDepth, root, 0)
return result
}

func process(result *int, maxDepth *int, root *TreeNode, depth int) {
if root == nil {
return
}
if root.Left == nil && root.Right == nil {
if depth > *maxDepth {
*maxDepth = depth
*result = root.Val
}
return
}
if root.Left != nil {
process(result, maxDepth, root.Left, depth+1)
}
if root.Right != nil {
process(result, maxDepth, root.Right, depth+1)
}
}
//leetcode submit region end(Prohibit modification and deletion)
