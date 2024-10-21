//给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历） 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：[[15,7],[9,20],[3]]
// 
//
// 示例 2： 
//
// 
//输入：root = [1]
//输出：[[1]]
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 2000] 内 
// -1000 <= Node.val <= 1000 
// 
//
// Related Topics 树 广度优先搜索 二叉树 👍 813 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    res := make([][]int, 0)
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        tmp, newQueue := []int{}, make([]*TreeNode, 0)
        for i:=0;i<len(queue);i++ {
            tmp = append(tmp, queue[i].Val)
            if queue[i].Left != nil {
                newQueue = append(newQueue, queue[i].Left)
            }
            if queue[i].Right != nil {
                newQueue = append(newQueue, queue[i].Right)
            }
        }
        queue = newQueue
        res = append(res, tmp)
    }
    i,j := 0, len(res)-1
    for i<j {
        res[i], res[j] = res[j], res[i]
        i++
        j--
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
