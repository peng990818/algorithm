//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。 
//
// 
//
// 示例1： 
//
// 
//
// 
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
// 
//
// 示例2： 
//
// 
//输入: root = [1,2,3]
//输出: [1,3]
// 
//
// 
//
// 提示： 
//
// 
// 二叉树的节点个数的范围是 [0,10⁴] 
// 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
//
// 
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 397 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
if root == nil {
return nil
}
queue := []*TreeNode{root}
result := make([]int, 0)
for len(queue) > 0 {
size := len(queue)
maxVal := math.MinInt32
for i := 0; i < size; i++ {
if queue[i].Val > maxVal {
maxVal = queue[i].Val
}
if queue[i].Left != nil {
queue = append(queue, queue[i].Left)
}
if queue[i].Right != nil {
queue = append(queue, queue[i].Right)
}
}
result = append(result, maxVal)
queue = queue[size:]
}
return result
}
//leetcode submit region end(Prohibit modification and deletion)
